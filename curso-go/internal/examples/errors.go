package examples

import (
	"errors"
	"fmt"
)

func Divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("divisor can't be zero")
	}
	return dividend / divisor, nil
}

func TryErrors() {

	result, err := Divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}

type ApiError struct {
	Code    int
	Message string
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("Error code: %v, Message: %s", e.Code, e.Message)
}
