package utility

import (
	"fmt"
	"strings"
)

type CustomError struct {
	Errors []ErrorPair
}

type ErrorPair struct {
	Field   string
	Message string
}

// Errorメソッドはerrorインターフェースのメソッドを実装します。
func (e *CustomError) Error() string {
	var errorStrings []string
	for _, err := range e.Errors {
		errorStrings = append(errorStrings, fmt.Sprintf("%s: %s", err.Field, err.Message))
	}
	if len(errorStrings) == 1 {
		return errorStrings[0] + ";"
	}
	return strings.Join(errorStrings, "; ")
}
