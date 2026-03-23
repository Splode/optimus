package image

import (
	"context"
	"encoding/json"
	"fmt"
	"optimus/backend/config"
	"optimus/backend/stat"
	"os"
	"path/filepath"
	"runtime/debug"
	"strconv"
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

// SelectFileMeta holds UI-facing metadata returned after a file is selected and
// decoded. The heavy image data stays in Go; only what the frontend needs is
// returned across the IPC bridge.
type SelectFileMeta struct {
	ID       string `json:"id"`
	Filename string `json:"filename"`
	Name     string `json:"name"`
	Size     int64  `json:"size"`
}

// SelectFiles opens a native OS file dialog, reads and decodes the selected
// image files, adds them to the FileManager, and returns their metadata for
// the UI. This avoids the unreliable browser <input type="file"> path on macOS.
func (fm *FileManager) SelectFiles() ([]SelectFileMeta, error) {
	paths, err := runtime.OpenMultipleFilesDialog(fm.ctx, runtime.OpenDialogOptions{
		Filters: []runtime.FileFilter{
			{DisplayName: "Images (*.jpg, *.jpeg, *.png, *.webp)", Pattern: "*.jpg;*.jpeg;*.png;*.webp"},
		},
	})
	if err != nil || len(paths) == 0 {
		return nil, err
	}

	var results []SelectFileMeta
	for _, p := range paths {
		data, err := os.ReadFile(p)
		if err != nil {
			runtime.LogErrorf(fm.ctx, "failed to read file: %s: %v", p, err)
			continue
		}
		base := filepath.Base(p)
		extWithDot := filepath.Ext(base)
		ext := strings.ToLower(strings.TrimPrefix(extWithDot, "."))
		name := strings.TrimSuffix(base, extWithDot)
		size := int64(len(data))
		id := name + strconv.FormatInt(size, 10)

		mime := "image/" + ext
		if ext == "jpg" {
			mime = "image/jpeg"
		}

		file := &File{
			Data:     data,
			Ext:      ext,
			ID:       id,
			MimeType: mime,
			Name:     name,
			Size:     size,
		}
		if err := file.Decode(); err != nil {
			runtime.LogErrorf(fm.ctx, "failed to decode file %s: %v", p, err)
			continue
		}
		file.Data = nil // release raw bytes; image is now in file.Image
		fm.Files = append(fm.Files, file)
		runtime.LogInfof(fm.ctx, "added file to file manager via dialog: %s", file.Name)

		results = append(results, SelectFileMeta{
			ID:       id,
			Filename: base,
			Name:     name,
			Size:     size,
		})
	}
	return results, nil
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
