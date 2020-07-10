package config

import (
	"fmt"
	"github.com/wailsapp/wails"
	"os"
	"path"
)

// Config represents the application settings.
type Config struct {
	OutDir  string
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
		_ = fmt.Errorf("failed to get user directory: %v", err)
	}

	od := path.Join(ud, "optimus")

	if _, err := os.Stat(od); os.IsNotExist(err) {
		if err := os.Mkdir(od, 0666); err != nil {
			od = "./"
			_ = fmt.Errorf("failed to create default output directory: %v", err)
		}
	}

	c.OutDir = od
	return c
}

// SetOutDir opens a directory select dialog and sets the output directory to
// the chosen directory.
func (c *Config) SetOutDir() string {
	dir := c.Runtime.Dialog.SelectDirectory()
	c.OutDir = dir
	return c.OutDir
}
