package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
	"optimus/backend/config"
	"optimus/backend/image"
	"optimus/backend/stat"
)

func main() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:     1200,
		Height:    742,
		Title:     "Optimus",
		JS:        js,
		CSS:       css,
		Colour:    "#18181f",
		Resizable: true,
	})

	cfg := config.NewConfig()
	st := stat.NewStat()
	fm := image.NewFileManager(cfg, st)

	app.Bind(cfg)
	app.Bind(st)
	app.Bind(fm)
	_ = app.Run()
}
