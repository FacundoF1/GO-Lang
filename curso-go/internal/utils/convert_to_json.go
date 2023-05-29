package convert_to_json

import (
	"encoding/json"
	"fmt"
)

func ConvertToJson(data map[string]interface{}) interface{} {
	response, _ := json.MarshalIndent(data, "", "\t")
	fmt.Printf("A map with values: %v\n", response)
	return response
}