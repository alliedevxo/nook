package main

import (
	"embed"
	"fmt"
	"log"
	nookapp "nook/internal/app"
	nookdb "nook/internal/db"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:client/dist
var assets embed.FS

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	store, err := nookdb.Open()

	if err != nil {
		return fmt.Errorf("open notebook db: %w", err)
	}
	defer store.Close()

	app := nookapp.New(store)

	return wails.Run(&options.App{
		Title: "Nook",
		Width: 1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.Startup,
		Bind: []interface{}{
			app,
		},
	})
}