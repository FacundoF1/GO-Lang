package convert_to_json

import (
	"encoding/json"
)

func ConvertToJson(data any) string {
	bytes, _ := json.MarshalIndent(data, "", "\t")
	return string(bytes)
}