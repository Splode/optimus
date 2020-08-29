package image

import (
	"encoding/json"
	"fmt"
	"github.com/wailsapp/wails"
	"optimus/lib/config"
	"optimus/lib/stat"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

type FileManager struct {
	Files []*File

	Runtime *wails.Runtime
	Logger  *wails.CustomLogger

	config *config.Config
	stats  *stat.Stat
}

// NewFileManager creates a new FileManager.
func NewFileManager(c *config.Config, s *stat.Stat) *FileManager {
	return &FileManager{
		config: c,
		stats:  s,
	}
}

// WailsInit performs setup when Wails is ready.
func (fm *FileManager) WailsInit(runtime *wails.Runtime) error {
	fm.Runtime = runtime
	fm.Logger = fm.Runtime.Log.New("FileManager")
	fm.Logger.Info("FileManager initialized...")
	return nil
}

// HandleFile processes a file from the client.
func (fm *FileManager) HandleFile(fileJson string) (err error) {
	file := &File{Runtime: fm.Runtime}
	if err := json.Unmarshal([]byte(fileJson), &file); err != nil {
		return err
	}

	if err = file.Decode(); err != nil {
		return err
	}
	fm.Files = append(fm.Files, file)
	fm.Logger.Infof("added file to file manager: %s", file.Name)

	return nil
}

// Clear removes the files in the FileManager.
func (fm *FileManager) Clear() {
	fm.Files = nil
	debug.FreeOSMemory()
}

// Convert runs the conversion on all files in the FileManager.
func (fm *FileManager) Convert() (errs []error) {
	var wg sync.WaitGroup
	wg.Add(fm.countUnconverted())

	c := 0
	var b int64
	t := time.Now().UnixNano()
	for _, file := range fm.Files {
		file := file
		if !file.IsConverted {
			go func(wg *sync.WaitGroup) {
				err := file.Write(fm.config)
				if err != nil {
					fm.Logger.Errorf("failed to convert file: %s, %v", file.ID, err)
					fm.Runtime.Events.Emit("notify", map[string]interface{}{
						"msg":  fmt.Sprintf("Failed to convert file: %s, %s", file.Name, err.Error()),
						"type": "warn",
					})
					errs = append(errs, fmt.Errorf("failed to convert file: %s", file.Name))
				} else {
					fm.Logger.Info(fmt.Sprintf("converted file: %s", file.Name))
					s, err := file.GetConvertedSize()
					if err != nil {
						fm.Logger.Errorf("failed to read converted file size: %v", err)
					}
					fm.Runtime.Events.Emit("conversion:complete", map[string]interface{}{
						"id": file.ID,
						// TODO: standardize this path conversion
						"path": strings.Replace(file.ConvertedFile, "\\", "/", -1),
						"size": s,
					})
					c++
					s, err = file.GetSavings()
					if err != nil {
						fm.Logger.Errorf("failed to get file conversion savings: %v", err)
					}
					b += s
				}
				wg.Done()
			}(&wg)
		}
	}

	wg.Wait()
	nt := (time.Now().UnixNano() - t) / 1000000
	fm.stats.SetImageCount(c)
	fm.stats.SetByteCount(b)
	fm.stats.SetTimeCount(nt)
	fm.Runtime.Events.Emit("conversion:stat", map[string]interface{}{
		"count":   c,
		"resizes": c * len(fm.config.App.Sizes),
		"savings": b,
		"time":    nt,
	})
	fm.Clear()
	return errs
}

// OpenFile opens the file at the given filepath using the file's native file
// application.
func (fm *FileManager) OpenFile(p string) error {
	if err := fm.Runtime.Browser.OpenFile(p); err != nil {
		fm.Logger.Errorf("failed to open file %s: %v", p, err)
		return err
	}
	return nil
}

// countUnconverted returns the number of files in the FileManager that haven't
// been converted.
func (fm *FileManager) countUnconverted() int {
	c := 0
	for _, file := range fm.Files {
		if !file.IsConverted {
			c++
		}
	}
	return c
}
