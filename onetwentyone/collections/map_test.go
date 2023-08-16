package collections_test

import (
	"maps"
	"testing"
)

func TestEqualMaps(t *testing.T) {
	t.Parallel()
	// Given
	m1 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	m2 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	want := true
	// When
	got := maps.Equal(m1, m2)
	// Then
	if want != got {
		t.Errorf("want: %t, but got: %t", want, got)
	}
}

func TestMapClone(t *testing.T) {
	t.Parallel()
	// Given
	m1 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	// When
	got := maps.Clone(m1)
	// Then
	if !maps.Equal(m1, got) {
		t.Errorf("map was not cloned")
	}
}

func TestCopyClone(t *testing.T) {
	t.Parallel()
	// Given
	m1 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	got := make(map[string]int, len(m1))
	// When
	maps.Copy(got, m1)
	// Then
	if !maps.Equal(m1, got) {
		t.Errorf("map was not cloned")
	}
}
