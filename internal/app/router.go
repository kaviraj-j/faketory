package app

import "github.com/gin-gonic/gin"

func (app *App) setupRoutes(router *gin.Engine) {

	router.GET("/posts", app.mockDataHandler.GetPosts)
	router.GET("/posts/:id", app.mockDataHandler.GetPost)

	router.GET("/users/:id", app.mockDataHandler.GetUser)
	router.GET("/users", app.mockDataHandler.GetUsers)

}
