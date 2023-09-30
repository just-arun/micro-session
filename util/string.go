package util

import (
	"encoding/json"

	"github.com/google/uuid"
)

func NewUUID() string {
	return uuid.New().String()
}

func JsonToMap(jsonStr string) map[string]interface{} {
	result := make(map[string]interface{})
	json.Unmarshal([]byte(jsonStr), &result)
	return result
}
