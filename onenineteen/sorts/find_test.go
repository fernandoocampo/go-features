package sorts_test

import (
	"sort"
	"testing"
)

func TestSortFind(t *testing.T) {
	// Given
	values := []int{9, 3, 2, 7, 19, 27, 34, 23, 78, 0, 1}
	sort.Ints(values)
	searchedValue := 19
	// When
	got, exists := sort.Find(len(values), func(i int) int {
		if values[i] == searchedValue {
			return 0
		}
		if searchedValue < values[i] {
			return -1
		}
		return 1
	})
	// Then
	if !exists {
		t.Errorf("want %d exists but it does not", searchedValue)
	}
	if values[got] != searchedValue {
		t.Errorf("want: %d, but got: %d", searchedValue, got)
	}
}
