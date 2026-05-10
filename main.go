package main

import (
	"embed"
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
	db, err := nookdb.Open()

	if err != nil {
		log.Fatalf("open notebook db: %v", err)
	}
	defer db.Close()

	app := nookapp.New(db)

	err = wails.Run(&options.App{
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
	if err != nil {
		log.Fatalf("wails run: %v", err)
	}
}