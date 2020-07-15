package config

import (
	"encoding/json"
	"fmt"
	"github.com/wailsapp/wails"
	"optimus/lib/localstore"
	"os"
	"path"
	"path/filepath"
)

const filename = "conf.json"

// App represents application persistent configuration values.
type App struct {
	OutDir string `json:"outDir"`
	Target string `json:"target"`
}

// Config represents the application settings.
type Config struct {
	App        *App
	LocalStore *localstore.LocalStore
	Runtime    *wails.Runtime
	Logger     *wails.CustomLogger
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
	c.LocalStore = localstore.NewLocalStore()

	a, err := c.LocalStore.Load(filename)
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
		"outDir": c.App.OutDir,
		"target": c.App.Target,
	}
}

// OpenOutputDir opens the output directory using the native system browser.
func (c *Config) OpenOutputDir() error {
	if err := c.Runtime.Browser.OpenURL(c.App.OutDir); err != nil {
		return err
	}
	return nil
}

// SetOutDir opens a directory select dialog and sets the output directory to
// the chosen directory.
func (c *Config) SetOutDir() string {
	dir := c.Runtime.Dialog.SelectDirectory()
	c.App.OutDir = dir
	c.Logger.Infof("set output directory: %s", dir)
	if err := c.Store(); err != nil {
		c.Logger.Errorf("failed to store config: %v", err)
	}
	return c.App.OutDir
}

// SetTarget sets the conversion output target.
func (c *Config) SetTarget(t string) error {
	// TODO check if valid target
	c.App.Target = t
	c.Logger.Infof("set conversion target: %s", t)
	if err := c.Store(); err != nil {
		return err
	}
	return nil
}

// Store stores the configuration state to the file system.
func (c *Config) Store() error {
	js, err := json.Marshal(c.GetAppConfig())
	if err != nil {
		c.Logger.Errorf("failed to marshal config: %v", err)
		return err
	}
	if err := c.LocalStore.Store(js, filename); err != nil {
		c.Logger.Errorf("failed to store config: %v", err)
		return err
	}
	return nil
}

// defaults returns the application configuration defaults.
func defaults() (*App, error) {
	a := &App{Target: "webp"}
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
