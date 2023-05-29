package file_system

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	for i := 0; i < 10; i++ {
        result := ReadFile("../../Electric_Vehicle_Population_Data.csv")
		assert.Equal(t, 130444, len(result))
    }
}

func TestReadFileBenchmark(t *testing.T) {
	for i := 0; i < 10; i++ {
      	ReadFile("../../Electric_Vehicle_Population_Data.csv")
    }
}