package service

import (
	"encoding/json"
	"log"
	"os"

	"github.com/kaviraj-j/faketory/internal/types"
)

func getMockData() (map[int]types.User, map[int]types.Post) {
	filePath := getFilePath()
	var users []types.User
	var posts []types.Post
	err := readData(filePath+"/users.json", &users)
	if err != nil {
		log.Fatal(err)
	}
	err = readData(filePath+"/posts.json", &posts)
	if err != nil {
		log.Fatal(err)
	}
	return convertToMap(users), convertToMap(posts)
}

func convertToMap[T any](arr []T) map[int]T {
	m := make(map[int]T)
	for _, v := range arr {
		switch t := any(v).(type) {
		case types.User:
			m[t.ID] = v
		case types.Post:
			m[t.ID] = v
		}
	}
	return m
}

func readData(path string, v any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func getFilePath() string {
	return os.Getenv("DATA_PATH")
}
