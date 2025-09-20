package service

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/kaviraj-j/faketory/internal/types"
)

func getMockData() ([]types.User, []types.Post, []types.Todo, []types.Album, []types.Comment) {
	filePath := getFilePath()

	// Data path mapping for generic loading
	dataPaths := map[string]interface{}{
		"users":    &[]types.User{},
		"posts":    &[]types.Post{},
		"todos":    &[]types.Todo{},
		"albums":   &[]types.Album{},
		"comments": &[]types.Comment{},
	}

	// Load all data files
	for dataType, dataPtr := range dataPaths {
		err := readData(filePath+"/"+dataType+".json", dataPtr)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Extract the loaded data
	users := *dataPaths["users"].(*[]types.User)
	posts := *dataPaths["posts"].(*[]types.Post)
	todos := *dataPaths["todos"].(*[]types.Todo)
	albums := *dataPaths["albums"].(*[]types.Album)
	comments := *dataPaths["comments"].(*[]types.Comment)

	return users, posts, todos, albums, comments
}

func readData(path string, v any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func getFilePath() string {
	return "data"
}

// ParseQueryParams parses and validates query parameters
func ParseQueryParams(countStr, userIDStr, postIDStr, idStr string) types.QueryParams {
	params := types.QueryParams{}

	if countStr != "" {
		if count, err := strconv.Atoi(countStr); err == nil && count > 0 {
			params.Count = count
		}
	}

	if userIDStr != "" {
		if userID, err := strconv.Atoi(userIDStr); err == nil && userID > 0 {
			params.UserID = userID
		}
	}

	if postIDStr != "" {
		if postID, err := strconv.Atoi(postIDStr); err == nil && postID > 0 {
			params.PostID = postID
		}
	}

	if idStr != "" {
		if id, err := strconv.Atoi(idStr); err == nil && id > 0 {
			params.ID = id
		}
	}

	return params
}
