package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kaviraj-j/faketory/internal/handlers"
	"github.com/kaviraj-j/faketory/internal/service"
	"github.com/kaviraj-j/faketory/internal/types"
)

type App struct {
	config          types.Config
	mockDataHandler *handlers.MockDataHandler
}

func NewApp(config types.Config) *App {
	mockDataService := service.NewMockDataService()
	handler := handlers.NewMockDataHandler(mockDataService)
	return &App{
		config:          config,
		mockDataHandler: handler,
	}
}

func (app *App) Run() {
	router := gin.Default()
	app.setupRoutes(router)
	listenUrl := fmt.Sprintf("%s:%s", app.config.AppUrl, app.config.AppPort)

	router.Run(listenUrl)
}
