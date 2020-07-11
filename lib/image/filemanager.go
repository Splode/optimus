package image

import (
	"encoding/json"
	"fmt"
	"github.com/wailsapp/wails"
	"optimus/lib/config"
	"sync"
)

type FileManager struct {
	Config  *config.Config
	Files   []*File
	OutDir  string
	Runtime *wails.Runtime
	Logger  *wails.CustomLogger
}

// NewFileManager creates a new FileManager.
func NewFileManager(c *config.Config) *FileManager {
	return &FileManager{
		Config: c,
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
	file := &File{}
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

// Convert runs the conversion on all files in the FileManager.
func (fm *FileManager) Convert() (errs []error) {
	var wg sync.WaitGroup
	wg.Add(fm.CountUnconverted())

	for _, file := range fm.Files {
		if !file.IsConverted {
			file := file
			go func(wg *sync.WaitGroup) {
				err := file.Write(fm.Config.OutDir)
				if err != nil {
					fm.Logger.Errorf("failed to convert file: %s, %v", file.ID, err)
					errs = append(errs, fmt.Errorf("failed to convert file: %s", file.Name))
				} else {
					file.IsConverted = true
					fm.Logger.Info(fmt.Sprintf("converted file: %s", file.Name))
					s, err := file.GetConvertedSize()
					if err != nil {
						fm.Logger.Errorf("failed to read converted file size: %v", err)
					}
					fm.Runtime.Events.Emit("conversion:complete", map[string]interface{}{
						"id":   file.ID,
						"size": s,
					})
				}
				wg.Done()
			}(&wg)
		}
	}

	wg.Wait()
	return errs
}

// CountUnconverted returns the number of files in the FileManager that haven't
// been converted.
func (fm *FileManager) CountUnconverted() int {
	c := 0
	for _, file := range fm.Files {
		if !file.IsConverted {
			c++
		}
	}
	return c
}
