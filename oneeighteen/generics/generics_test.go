package generics_test

import (
	"testing"

	"github.com/fernandoocampo/go-features/oneeighteen/generics"
)

func TestPrintSlice(t *testing.T) {
	t.Run("string type", func(t *testing.T) {
		// Given
		want := "drila, michael, ocho, uno, nino, ivana"
		given := []string{"drila", "michael", "ocho", "uno", "nino", "ivana"}
		// When
		got := generics.PrintSlice[string](given)
		if want != got {
			t.Errorf("want: %s, but got: %s", want, got)
		}
	})
	t.Run("int type", func(t *testing.T) {
		// Given
		want := "1, 2, 3, 4, 5, 6"
		given := []int{1, 2, 3, 4, 5, 6}
		// When
		got := generics.PrintSlice(given)
		if want != got {
			t.Errorf("want: %s, but got: %s", want, got)
		}
	})
}

func TestLastItem(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		// Given
		want := "5"
		given := generics.List[string]{"1", "2", "3", "4", "5"}
		// When
		got, err := given.LastItem()
		// Then
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if want != got {
			t.Errorf("want: %s, but got: %s", want, got)
		}
	})
	t.Run("int", func(t *testing.T) {
		// Given
		want := 5
		given := generics.List[int]([]int{1, 2, 3, 4, 5})
		// When
		got, err := given.LastItem()
		// Then
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if want != got {
			t.Errorf("want: %d, but got: %d", want, got)
		}
	})
}
