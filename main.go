package main

import (
	"context"
	"embed"
	"optimus/backend/config"
	"optimus/backend/image"
	"optimus/backend/stat"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	cfg := config.NewConfig()
	st := stat.NewStat()
	fm := image.NewFileManager(cfg, st)

	err := wails.Run(&options.App{
		Title:  "Optimus",
		Width:  1200,
		Height: 742,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 24, G: 24, B: 31, A: 255},
		OnStartup: func(ctx context.Context) {
			cfg.Startup(ctx)
			st.Startup(ctx)
			fm.Startup(ctx)
		},
		Bind: []interface{}{cfg, st, fm},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
