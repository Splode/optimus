package config

import (
	"encoding/json"
	"fmt"
	"github.com/wailsapp/wails"
	"optimus/lib/jpeg"
	"optimus/lib/localstore"
	"optimus/lib/png"
	"optimus/lib/webp"
	"os"
	"path"
	"path/filepath"
)

const filename = "conf.json"

// App represents application persistent configuration values.
type App struct {
	OutDir  string        `json:"outDir"`
	Target  string        `json:"target"`
	Prefix  string        `json:"prefix"`
	Suffix  string        `json:"suffix"`
	JpegOpt *jpeg.Options `json:"jpegOpt"`
	PngOpt  *png.Options  `json:"pngOpt"`
	WebpOpt *webp.Options `json:"webpOpt"`
}

// Config represents the application settings.
type Config struct {
	App        *App
	Runtime    *wails.Runtime
	Logger     *wails.CustomLogger
	localStore *localstore.LocalStore
}

// WailsInit performs setup when Wails is ready.
func (c *Config) WailsInit(runtime *wails.Runtime) error {
	c.Runtime = runtime
	c.Logger = c.Runtime.Log.New("Config")
	c.Logger.Info("Config initialized...")
	return nil
}

// NewConfig returns a new instance of Config.
func NewConfig() *Config {
	c := &Config{}
	c.localStore = localstore.NewLocalStore()

	a, err := c.localStore.Load(filename)
	if err != nil {
		c.App, _ = defaults()
	}
	if err := json.Unmarshal(a, &c.App); err != nil {
		fmt.Printf("error")
	}
	return c
}

// GetAppConfig returns the application configuration.
func (c *Config) GetAppConfig() map[string]interface{} {
	return map[string]interface{}{
		"outDir":  c.App.OutDir,
		"target":  c.App.Target,
		"prefix":  c.App.Prefix,
		"suffix":  c.App.Suffix,
		"jpegOpt": c.App.JpegOpt,
		"pngOpt":  c.App.PngOpt,
		"webpOpt": c.App.WebpOpt,
	}
}

// OpenOutputDir opens the output directory using the native system browser.
func (c *Config) OpenOutputDir() error {
	if err := c.Runtime.Browser.OpenURL(c.App.OutDir); err != nil {
		return err
	}
	return nil
}

// SetConfig sets and stores the given configuration.
func (c *Config) SetConfig(cfg string) error {
	a := &App{}
	if err := json.Unmarshal([]byte(cfg), &a); err != nil {
		c.Logger.Errorf("failed to unmarshal config: %v", err)
		return err
	}
	c.App = a
	if err := c.store(); err != nil {
		c.Logger.Errorf("failed to store config: %v", err)
		return err
	}
	return nil
}

// SetOutDir opens a directory select dialog and sets the output directory to
// the chosen directory.
func (c *Config) SetOutDir() string {
	dir := c.Runtime.Dialog.SelectDirectory()
	if dir != "" {
		c.App.OutDir = dir
		c.Logger.Infof("set output directory: %s", dir)
		if err := c.store(); err != nil {
			c.Logger.Errorf("failed to store config: %v", err)
		}
	}
	return c.App.OutDir
}

// defaults returns the application configuration defaults.
func defaults() (*App, error) {
	a := &App{
		Target:  "webp",
		JpegOpt: &jpeg.Options{Quality: 80},
		PngOpt:  &png.Options{Quality: 80},
		WebpOpt: &webp.Options{Lossless: false, Quality: 80},
	}
	ud, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("failed to get user directory: %v", err)
		return nil, err
	}

	od := path.Join(ud, "Optimus")
	cp := filepath.Clean(od)

	if _, err := os.Stat(od); os.IsNotExist(err) {
		if err := os.Mkdir(od, 0777); err != nil {
			od = "./"
			fmt.Printf("failed to create default output directory: %v", err)
			return nil, err
		}
	}
	a.OutDir = cp
	return a, nil
}

// store stores the configuration state to the file system.
func (c *Config) store() error {
	js, err := json.Marshal(c.GetAppConfig())
	if err != nil {
		c.Logger.Errorf("failed to marshal config: %v", err)
		return err
	}
	if err := c.localStore.Store(js, filename); err != nil {
		c.Logger.Errorf("failed to store config: %v", err)
		return err
	}
	return nil
}
