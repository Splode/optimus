package config

import (
	"fmt"
	"github.com/wailsapp/wails"
	"os"
	"path"
	"path/filepath"
)

// Config represents the application settings.
type Config struct {
	OutDir  string
	Target  string
	Runtime *wails.Runtime
	Logger  *wails.CustomLogger
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
	ud, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("failed to get user directory: %v", err)
	}

	od := path.Join(ud, "optimus")
	cp := filepath.Clean(od)

	if _, err := os.Stat(od); os.IsNotExist(err) {
		if err := os.Mkdir(od, 0777); err != nil {
			od = "./"
			fmt.Printf("failed to create default output directory: %v", err)
		}
	}

	c.OutDir = cp

	c.Target = "webp"
	return c
}

// GetAppConfig returns the application configuration.
func (c *Config) GetAppConfig() map[string]interface{} {
	return map[string]interface{}{
		"outDir": c.OutDir,
		"target": c.Target,
	}
}

// OpenOutputDir opens the output directory using the native system browser.
func (c *Config) OpenOutputDir() error {
	if err := c.Runtime.Browser.OpenURL(c.OutDir); err != nil {
		return err
	}
	return nil
}

// SetOutDir opens a directory select dialog and sets the output directory to
// the chosen directory.
func (c *Config) SetOutDir() string {
	dir := c.Runtime.Dialog.SelectDirectory()
	c.OutDir = dir
	c.Logger.Infof("set output directory: %s", dir)
	return c.OutDir
}

// SetTarget sets the conversion output target.
func (c *Config) SetTarget(t string) {
	// TODO check if valid target
	c.Target = t
	c.Logger.Infof("set conversion target: %s", t)
}
