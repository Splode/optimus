package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
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

	fm := image.NewFileManager()

	app.Bind(fm)
	_ = app.Run()
}
