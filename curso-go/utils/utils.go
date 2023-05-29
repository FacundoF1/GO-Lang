package utils

import (
	"encoding/json"
)

func ConvertToJsonInString(data any) string {
	bytes, _ := json.MarshalIndent(data, "", "\t")
	return string(bytes)
}