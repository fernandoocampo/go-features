package times_test

import (
	"testing"
	"time"
)

func TestDurationABS(t *testing.T) {
	// Given
	twoSecond := 2 * time.Second
	fiveSecond := 5 * time.Second
	result := twoSecond - fiveSecond
	want := 3 * time.Second
	// When
	got := result.Abs()
	// Then
	if result > 0 {
		t.Errorf("%s must be a positive number", result)
	}
	if got != want {
		t.Errorf("want: %s, but got: %s", want, got)
	}
}
