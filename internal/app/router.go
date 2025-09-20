package app

import "github.com/gin-gonic/gin"

func (app *App) setupRoutes(router *gin.Engine) {

	// Posts routes
	router.GET("/posts", app.mockDataHandler.GetPosts)
	router.GET("/posts/:id", app.mockDataHandler.GetPost)

	// Users routes
	router.GET("/users", app.mockDataHandler.GetUsers)
	router.GET("/users/:id", app.mockDataHandler.GetUser)

	// Todos routes
	router.GET("/todos", app.mockDataHandler.GetTodos)
	router.GET("/todos/:id", app.mockDataHandler.GetTodo)

	// Albums routes
	router.GET("/albums", app.mockDataHandler.GetAlbums)
	router.GET("/albums/:id", app.mockDataHandler.GetAlbum)

	// Comments routes
	router.GET("/comments", app.mockDataHandler.GetComments)
	router.GET("/comments/:id", app.mockDataHandler.GetComment)

}
