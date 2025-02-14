package main

import (
	"embed"
	"log"

	"todo-app/internal/app"
	"todo-app/internal/dal"
	"todo-app/internal/models"
	"todo-app/internal/service"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Todo{})

	todoRepo := dal.NewTodoRepo()
	todoService := service.NewTodoService(todoRepo)
	app := app.NewApp(todoService)

	err = wails.Run(&options.App{
		Title:  "todo-app",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 255},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
