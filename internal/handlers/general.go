package handlers

import (
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

func GetPosts(ctx *gin.Context) {

}
