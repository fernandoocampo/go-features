package generics

import (
	"bytes"
	"errors"
	"fmt"
)

func PrintSlice[T any](values []T) string {
	if len(values) == 0 {
		return ""
	}

	var buffer bytes.Buffer
	for _, v := range values {
		buffer.WriteString(fmt.Sprintf("%v", v))
		buffer.WriteString(", ")
	}
	result := buffer.String()
	return result[:len(result)-2]
}

type List[T any] []T

func (v List[T]) LastItem() (T, error) {
	var value T

	if len(v) == 0 {
		return value, errors.New("empty")
	}
	return v[len(v)-1], nil
}
