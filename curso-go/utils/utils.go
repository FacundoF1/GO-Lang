package utils

import (
	"encoding/json"
)

func ConvertToJsonInString(data any) string {
	bytes, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
        panic(err)
    }
	return string(bytes)
}