package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
	"optimus/lib/config"
	"optimus/lib/image"
)

func main() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "Optimus",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	cfg := config.NewConfig()
	fm := image.NewFileManager(cfg)

	app.Bind(cfg)
	app.Bind(fm)
	_ = app.Run()
}
