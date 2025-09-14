package main

import (
	"context"
	"embed"
	"fmt"
	app "to-do-app/internal/app"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

const (
	dir  = "./logs"
	mode = "debug"
)

func main() {
	ctx := context.Background()

	application, err := app.NewApp(ctx, dir, mode)
	if err != nil {
		fmt.Println("failed to setup application:", err)
		return
	}

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "to-do-app",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        application.Startup,
		OnShutdown:       application.Shutdown,
		Bind: []interface{}{
			application.TaskHandler,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
