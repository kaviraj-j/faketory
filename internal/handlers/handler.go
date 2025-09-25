package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaviraj-j/faketory/internal/service"
	"github.com/kaviraj-j/faketory/internal/types"
)

type MockDataHandler struct {
	mockDataService *service.MockDataService
}

func NewMockDataHandler(s *service.MockDataService) *MockDataHandler {
	return &MockDataHandler{
		mockDataService: s,
	}
}

// Helper function to parse query parameters
func (h *MockDataHandler) parseQueryParams(ctx *gin.Context) types.QueryParams {
	return service.ParseQueryParams(
		ctx.Query("count"),
		ctx.Query("userId"),
		ctx.Query("postId"),
		ctx.Param("id"),
	)
}

// User handlers
func (h *MockDataHandler) GetUsers(ctx *gin.Context) {
	params := h.parseQueryParams(ctx)
	users := h.mockDataService.GetUsers(params)
	ctx.JSON(http.StatusOK, users)
}

func (h *MockDataHandler) GetUser(ctx *gin.Context) {
	params := h.parseQueryParams(ctx)
	user, found := h.mockDataService.GetUser(params.ID)
	if !found {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// Post handlers
func (h *MockDataHandler) GetPosts(ctx *gin.Context) {
	params := h.parseQueryParams(ctx)
	posts := h.mockDataService.GetPosts(params)
	ctx.JSON(http.StatusOK, posts)
}

func (h *MockDataHandler) GetPost(ctx *gin.Context) {
	params := h.parseQueryParams(ctx)
	post, found := h.mockDataService.GetPost(params.ID)
	if !found {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, post)
}

// Todo handlers
func (h *MockDataHandler) GetTodos(ctx *gin.Context) {
	params := h.parseQueryParams(ctx)
	todos := h.mockDataService.GetTodos(params)
	ctx.JSON(http.StatusOK, todos)
}

func (h *MockDataHandler) GetTodo(ctx *gin.Context) {
	params := h.parseQueryParams(ctx)
	todo, found := h.mockDataService.GetTodo(params.ID)
	if !found {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Todo not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, todo)
}

// Album handlers
func (h *MockDataHandler) GetAlbums(ctx *gin.Context) {
	params := h.parseQueryParams(ctx)
	albums := h.mockDataService.GetAlbums(params)
	ctx.JSON(http.StatusOK, albums)
}

func (h *MockDataHandler) GetAlbum(ctx *gin.Context) {
	params := h.parseQueryParams(ctx)
	album, found := h.mockDataService.GetAlbum(params.ID)
	if !found {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Album not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, album)
}

// Comment handlers
func (h *MockDataHandler) GetComments(ctx *gin.Context) {
	params := h.parseQueryParams(ctx)
	comments := h.mockDataService.GetComments(params)
	ctx.JSON(http.StatusOK, comments)
}

func (h *MockDataHandler) GetComment(ctx *gin.Context) {
	params := h.parseQueryParams(ctx)
	comment, found := h.mockDataService.GetComment(params.ID)
	if !found {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Comment not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, comment)
}

func (h *MockDataHandler) GenerateData(ctx *gin.Context) {
	var mockDataDetails types.MockDataBody
	if err := ctx.BindJSON(&mockDataDetails); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"type":    "error",
			"message": "error in parsing body",
		})
		return
	}
	data, err := h.mockDataService.GenerateData(mockDataDetails.Schema, mockDataDetails.Count)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"type":    "error",
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, data)
}
