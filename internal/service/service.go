package service

import (
	"github.com/kaviraj-j/faketory/internal/types"
)

type MockDataService struct {
	mockData types.MockData
}

func NewMockDataService() *MockDataService {
	users, posts, todos, albums, comments := getMockData()

	return &MockDataService{
		mockData: types.MockData{
			Users:    users,
			Posts:    posts,
			Todos:    todos,
			Albums:   albums,
			Comments: comments,
		},
	}
}

func findByID[T any](items []T, id int, getID func(T) int) (T, bool) {
	var zero T
	for _, item := range items {
		if getID(item) == id {
			return item, true
		}
	}
	return zero, false
}

func validateCount(count int) int {
	if count <= 0 {
		return 10 // default
	}
	if count > 100 {
		return 100 // max limit
	}
	return count
}

// User methods
func (s *MockDataService) GetUsers(params types.QueryParams) []types.User {
	count := validateCount(params.Count)
	users := s.mockData.Users

	if count >= len(users) {
		result := make([]types.User, len(users))
		copy(result, users)
		return result
	}

	result := make([]types.User, count)
	copy(result, users[:count])
	return result
}

func (s *MockDataService) GetUser(id int) (types.User, bool) {
	return findByID(s.mockData.Users, id, func(u types.User) int { return u.ID })
}

// Post methods
func (s *MockDataService) GetPosts(params types.QueryParams) []types.Post {
	count := validateCount(params.Count)
	posts := s.mockData.Posts

	if count >= len(posts) {
		result := make([]types.Post, len(posts))
		copy(result, posts)
		return result
	}

	result := make([]types.Post, count)
	copy(result, posts[:count])
	return result
}

func (s *MockDataService) GetPost(id int) (types.Post, bool) {
	return findByID(s.mockData.Posts, id, func(p types.Post) int { return p.ID })
}

// Todo methods
func (s *MockDataService) GetTodos(params types.QueryParams) []types.Todo {
	count := validateCount(params.Count)
	todos := s.mockData.Todos

	if count >= len(todos) {
		result := make([]types.Todo, len(todos))
		copy(result, todos)
		return result
	}

	result := make([]types.Todo, count)
	copy(result, todos[:count])
	return result
}

func (s *MockDataService) GetTodo(id int) (types.Todo, bool) {
	return findByID(s.mockData.Todos, id, func(t types.Todo) int { return t.ID })
}

// Album methods
func (s *MockDataService) GetAlbums(params types.QueryParams) []types.Album {
	count := validateCount(params.Count)
	albums := s.mockData.Albums

	if count >= len(albums) {
		result := make([]types.Album, len(albums))
		copy(result, albums)
		return result
	}

	result := make([]types.Album, count)
	copy(result, albums[:count])
	return result
}

func (s *MockDataService) GetAlbum(id int) (types.Album, bool) {
	return findByID(s.mockData.Albums, id, func(a types.Album) int { return a.ID })
}

// Comment methods
func (s *MockDataService) GetComments(params types.QueryParams) []types.Comment {
	count := validateCount(params.Count)
	comments := s.mockData.Comments

	if count >= len(comments) {
		result := make([]types.Comment, len(comments))
		copy(result, comments)
		return result
	}

	result := make([]types.Comment, count)
	copy(result, comments[:count])
	return result
}

func (s *MockDataService) GetComment(id int) (types.Comment, bool) {
	return findByID(s.mockData.Comments, id, func(c types.Comment) int { return c.ID })
}
