package service

import (
	"encoding/json"
	"math"
	"math/rand"
)

func (s *MockDataService) GenerateData(schemaData json.RawMessage, count int) (any, error) {
	var schema map[string]string
	if err := json.Unmarshal(schemaData, &schema); err != nil {
		return nil, err
	}

	results := make([]map[string]any, 0, count)

	for i := 0; i < count; i++ {
		record := make(map[string]any)
		for field, fieldType := range schema {
			switch fieldType {
			case "string":
				record[field] = s.generateString()
			case "int":
				record[field] = s.generateInteger()
			case "float":
				record[field] = s.generateFloat()
			case "bool":
				record[field] = s.generateBool()
			default:
				record[field] = "UNSUPPORTED_TYPE"
			}
		}
		results = append(results, record)
	}

	return results, nil
}

func (s *MockDataService) generateString() string {
	count := rand.Intn(20) + 5
	byteData := make([]byte, 0, count)
	for i := 0; i < count; i++ {
		char := rand.Intn(52)
		if char < 26 {
			byteData = append(byteData, byte('A'+char))
		} else {
			byteData = append(byteData, byte('a'+char-26))
		}
	}
	return string(byteData)
}

func (s *MockDataService) generateInteger() int {
	return rand.Intn(1000)
}

func (s *MockDataService) generateFloat() float64 {
	return math.Round(rand.Float64()*10000) / 100
}

func (s *MockDataService) generateBool() bool {
	return rand.Intn(2) == 1
}
