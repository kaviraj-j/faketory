package service_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kaviraj-j/faketory/internal/service"
)

func TestGenerateData(t *testing.T) {
	var schema json.RawMessage
	jsonString := `
	{
		"name": "string",
		"age": "int",
		"grade":  "float",
		"completed": "bool"
    }
	`
	schema = json.RawMessage(jsonString)
	mockService := service.NewMockDataService()
	data, err := mockService.GenerateData(schema, 1)
	if err != nil {
		t.Errorf("expected error %q, got %v", "expected message", err)
	}
	fmt.Println(data)
}
