package convert_to_json

import (
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestConvertToJson(t *testing.T) {
	out := map[string]interface{}{}
	out["a"] = 1
 	result := ConvertToJson(out)
	const resultest string = string(result)
	fmt.Printf("A map with values tes: %v\n", result)
	assert.Equal(t, "{\n\t\"a\": 1,\n}", resultest)
}