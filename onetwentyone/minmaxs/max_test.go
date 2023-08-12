package minmaxs_test

import (
	"testing"
)

func TestMaxInt(t *testing.T) {
	t.Parallel()
	// Given
	wValue := 1
	xValue := 3
	yValue := 5
	zValue := 0
	want := 5
	// When
	got := max(wValue, xValue, yValue, zValue)
	// Then
	if want != got {
		t.Errorf("want: %d, but got: %d", want, got)
	}
}
