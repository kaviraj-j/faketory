package types

type Config struct {
	AppUrl  string
	AppPort string
}

type MockData struct {
	Users map[int]User
	Posts map[int]Post
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type Post struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
