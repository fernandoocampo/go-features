package collections_test

import (
	"testing"
)

func TestClearSlice(t *testing.T) {
	t.Parallel()
	// Given
	values := []int{1, 2, 3, 4, 9, 10}
	want := []int{0, 0, 0, 0, 0, 0}
	// When
	clear(values)
	// Then
	for index := range values {
		if want[index] != values[index] {
			t.Fatalf("want: %v, but got %v", want, values)
		}
	}
}

func TestClearMap(t *testing.T) {
	t.Parallel()
	// Given
	values := map[string]int{"one": 1, "two": 2, "three": 3}
	want := map[string]int{"one": 0, "two": 0, "three": 0}
	// When
	clear(values)
	// Then
	for key := range values {
		if want[key] != values[key] {
			t.Fatalf("want: %v, but got %v", want, values)
		}
	}
}

func TestClearSliceWithStruct(t *testing.T) {
	t.Parallel()
	// Given
	type dog struct {
		name string
		age  int
	}
	values := []dog{
		{name: "drila", age: 7},
		{name: "michael", age: 5},
		{name: "ivana", age: 8},
	}
	want := []dog{
		{name: "", age: 0},
		{name: "", age: 0},
		{name: "", age: 0},
	}
	// When
	clear(values)
	// Then
	for index := range values {
		if want[index] != values[index] {
			t.Fatalf("want: %v, but got %v", want, values)
		}
	}
}
