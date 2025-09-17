package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kaviraj-j/faketory/internal/types"
)

type App struct {
	config types.Config
	data   types.MockData
}

func NewApp(config types.Config) *App {
	return &App{
		config: config,
	}
}

func (app *App) Run() {
	router := gin.Default()
	app.setupRoutes(router)
	listenUrl := fmt.Sprintf("%s:%s", app.config.AppUrl, app.config.AppPort)

	router.Run(listenUrl)
}
