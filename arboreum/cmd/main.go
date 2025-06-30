package main

import (
	"arboreum/internal/service"
	"context"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

var assets embed.FS

func twin() {
	ctx := context.Background()
	service := service.NewTwinService(ctx)
	response, err := service.ProcessPrompt()
	if err != nil {
		panic(err)
	}
	println("Resposta do modelo:", response)
}

func main() {
	app := NewApp()
	err := wails.Run(&options.App{
		Title:  "arboreum",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
