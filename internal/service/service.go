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

// User methods
func (s *MockDataService) GetUsers(params types.QueryParams) []types.User {
	users := s.mockData.Users
	return users
}

func (s *MockDataService) GetUser(id int) (types.User, bool) {
	return findByID(s.mockData.Users, id, func(u types.User) int { return u.ID })
}

// Post methods
func (s *MockDataService) GetPosts(params types.QueryParams) []types.Post {
	posts := s.mockData.Posts
	return posts
}

func (s *MockDataService) GetPost(id int) (types.Post, bool) {
	return findByID(s.mockData.Posts, id, func(p types.Post) int { return p.ID })
}

// Todo methods
func (s *MockDataService) GetTodos(params types.QueryParams) []types.Todo {
	todos := s.mockData.Todos
	return todos
}

func (s *MockDataService) GetTodo(id int) (types.Todo, bool) {
	return findByID(s.mockData.Todos, id, func(t types.Todo) int { return t.ID })
}

// Album methods
func (s *MockDataService) GetAlbums(params types.QueryParams) []types.Album {
	albums := s.mockData.Albums
	return albums
}

func (s *MockDataService) GetAlbum(id int) (types.Album, bool) {
	return findByID(s.mockData.Albums, id, func(a types.Album) int { return a.ID })
}

// Comment methods
func (s *MockDataService) GetComments(params types.QueryParams) []types.Comment {
	comments := s.mockData.Comments
	return comments
}

func (s *MockDataService) GetComment(id int) (types.Comment, bool) {
	return findByID(s.mockData.Comments, id, func(c types.Comment) int { return c.ID })
}
