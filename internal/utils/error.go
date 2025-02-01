package utils

import "fmt"

type CustomError struct {
	ErrorCode int
	Err       error
}

func (err *CustomError) Error() string {
	return fmt.Sprintf("Error: %v, StatusCode: %d", err.Err, err.ErrorCode)
}
