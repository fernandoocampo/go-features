package minmaxs_test

import (
	"testing"
)

func TestMinInt(t *testing.T) {
	t.Parallel()
	// Given
	xValue := 3
	yValue := 5
	want := 3
	// When
	got := min(xValue, yValue)
	// Then
	if want != got {
		t.Errorf("want: %d, but got: %d", want, got)
	}
}

func TestMinMultipleValues(t *testing.T) {
	t.Parallel()
	// Given
	wValue := 1
	xValue := 3
	yValue := 5
	zValue := 0
	want := 0
	// When
	got := min(wValue, xValue, yValue, zValue)
	// Then
	if want != got {
		t.Errorf("want: %d, but got: %d", want, got)
	}
}

func TestMinString(t *testing.T) {
	t.Parallel()
	// Given
	xValue := "a"
	yValue := "b"
	want := "a"
	// When
	got := min(xValue, yValue)
	// Then
	if want != got {
		t.Errorf("want: %s, but got: %s", want, got)
	}
}
