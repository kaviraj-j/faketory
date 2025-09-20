package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kaviraj-j/faketory/internal/service"
)

type MockDataHandler struct {
	mockDataService *service.MockDataService
}

func NewMockDataHandler(s *service.MockDataService) *MockDataHandler {
	return &MockDataHandler{
		mockDataService: s,
	}
}

func (h *MockDataHandler) GetPosts(ctx *gin.Context) {
	countStr := ctx.Query("count")
	count, _ := strconv.Atoi(countStr)
	posts, _ := h.mockDataService.GetPosts(count)
	ctx.JSON(http.StatusOK, gin.H{
		"data": posts,
	})
}

func (h *MockDataHandler) GetPost(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	post, _ := h.mockDataService.GetPost(id)
	ctx.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}

func (h *MockDataHandler) GetUsers(ctx *gin.Context) {
	countStr := ctx.Query("count")
	count, _ := strconv.Atoi(countStr)
	users, _ := h.mockDataService.GetUsers(count)
	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (h *MockDataHandler) GetUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	user, _ := h.mockDataService.GetUser(id)
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
