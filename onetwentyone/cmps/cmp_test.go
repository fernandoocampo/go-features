package cmps_test

import (
	"cmp"
	"testing"
)

func TestCompare(t *testing.T) {
	t.Parallel()
	// Given
	tests := map[string]struct {
		x    uint8
		y    uint8
		want int
	}{
		"less_than": {
			x:    3,
			y:    10,
			want: -1,
		},
		"equals": {
			x:    13,
			y:    13,
			want: 0,
		},
		"greater_than": {
			x:    31,
			y:    10,
			want: 1,
		},
	}
	// When
	for name, data := range tests {
		testName := name
		testData := data
		t.Run(testName, func(tt *testing.T) {
			tt.Parallel()
			got := cmp.Compare(testData.x, testData.y)
			// Then
			if testData.want != got {
				tt.Errorf("want: %d, but got: %d", testData.want, got)
			}
		})
	}
}

func TestLess(t *testing.T) {
	// Given
	x := "car"
	y := "motorcycle"
	want := true
	// When
	got := cmp.Less(x, y)
	// Then
	if want != got {
		t.Errorf("want: %t, but got: %t", want, got)
	}
}
