package types

type Config struct {
	AppUrl  string
	AppPort string
}

type MockData struct {
	Users    []User
	Posts    []Post
	Todos    []Todo
	Albums   []Album
	Comments []Comment
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

type Todo struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Album struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
}

type Comment struct {
	ID     int    `json:"id"`
	PostID int    `json:"postId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

// Query parameters for filtering
type QueryParams struct {
	Count  int
	UserID int
	PostID int
	ID     int
}

// Data path mapping for generic retrieval
type DataPath struct {
	Path     string
	DataType string
}

type MockDataBody struct {
	DataKey string `json:"dataKey"`
	Count   string `json:"count"`
	Schema  string `json:"schema"`
}
