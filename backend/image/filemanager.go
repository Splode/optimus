package image

import (
	"context"
	"encoding/json"
	"fmt"
	"optimus/backend/config"
	"optimus/backend/stat"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// FileManager handles collections of Files for conversion.
type FileManager struct {
	Files []*File

	ctx    context.Context
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

// Startup stores the context when the app starts.
func (fm *FileManager) Startup(ctx context.Context) {
	fm.ctx = ctx
	runtime.LogInfo(ctx, "FileManager initialized...")
}

// HandleFile processes a file from the client.
func (fm *FileManager) HandleFile(fileJson string) (err error) {
	file := &File{}
	if err = json.Unmarshal([]byte(fileJson), &file); err != nil {
		return err
	}

	if err = file.Decode(); err != nil {
		return err
	}
	fm.Files = append(fm.Files, file)
	runtime.LogInfof(fm.ctx, "added file to file manager: %s", file.Name)

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
				err := file.Write(fm.config, fm.ctx)
				if err != nil {
					runtime.LogErrorf(fm.ctx, "failed to convert file: %s, %v", file.ID, err)
					runtime.EventsEmit(fm.ctx, "notify", map[string]interface{}{
						"msg":  fmt.Sprintf("Failed to convert file: %s, %s", file.Name, err.Error()),
						"type": "warn",
					})
					errs = append(errs, fmt.Errorf("failed to convert file: %s", file.Name))
				} else {
					runtime.LogInfof(fm.ctx, "converted file: %s", file.Name)
					s, err := file.GetConvertedSize()
					if err != nil {
						runtime.LogErrorf(fm.ctx, "failed to read converted file size: %v", err)
					}
					runtime.EventsEmit(fm.ctx, "conversion:complete", map[string]interface{}{
						"id": file.ID,
						// TODO: standardize this path conversion
						"path": strings.Replace(file.ConvertedFile, "\\", "/", -1),
						"size": s,
					})
					c++
					s, err = file.GetSavings()
					if err != nil {
						runtime.LogErrorf(fm.ctx, "failed to get file conversion savings: %v", err)
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
	runtime.EventsEmit(fm.ctx, "conversion:stat", map[string]interface{}{
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
	runtime.BrowserOpenURL(fm.ctx, p)
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
