package loops_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/fernandoocampo/go-features/onetwentythree/loops"
)

func TestRangeOverFunc(t *testing.T) {
	// Given
	values := []int{1, 2, 3, 4, 5, 6, 8, 9, 10}
	want := []int{2, 4, 6, 8, 10, 12, 16, 18, 20}
	// When
	got := loops.DoubleValues(values)
	// Then
	if !slices.Equal(want, got) {
		t.Errorf("want: %+v, but got: %+v, %d=%d", want, got, len(want), len(got))
	}
}

func TestNames(t *testing.T) {
	var myStructs loops.MyStructs = loops.MyStructs{
		loops.MyStruct{ID: 1, Name: "test1"},
		loops.MyStruct{ID: 3, Name: "test3"},
		loops.MyStruct{ID: 2, Name: "test2"},
	}
	for name := range myStructs.Names() {
		fmt.Printf("Name: %s\n", name)
	}
}
