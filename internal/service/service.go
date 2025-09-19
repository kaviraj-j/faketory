package service

import "github.com/kaviraj-j/faketory/internal/types"

type MockDataService struct {
	mockData types.MockData
}

func NewMockDataService() *MockDataService {
	return &MockDataService{}
}

func (s *MockDataService) GetUsers(count int) ([]types.User, error) {
	users := make([]types.User, 0, count)
	for _, u := range s.mockData.Users {
		users = append(users, u)
	}
	return users, nil
}

func (s *MockDataService) GetUser(id int) (types.User, error) {
	for _, u := range s.mockData.Users {
		if u.ID == id {
			return u, nil
		}
	}
	u := s.mockData.Users[0]
	u.ID = id
	return u, nil
}

func (s *MockDataService) GetPosts(count int) ([]types.Post, error) {
	posts := make([]types.Post, 0, count)
	for _, p := range s.mockData.Posts {
		posts = append(posts, p)
	}
	return posts, nil
}

func (s *MockDataService) GetPost(id int) (types.Post, error) {
	for _, p := range s.mockData.Posts {
		if p.ID == id {
			return p, nil
		}
	}
	p := s.mockData.Posts[0]
	p.ID = id
	return p, nil
}
