package collections_test

import (
	"slices"
	"sort"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	t.Parallel()
	// Given
	values := []string{
		"luigi", "mario",
		"peach", "bowser",
		"yoshi", "wario",
		"daisy", "rosalina",
		"toad", "lakitu",
		"koopa", "boo",
	}
	tests := map[string]struct {
		want int
		ok   bool
	}{
		"boo": {
			want: 0,
			ok:   true,
		},
		"bowser": {
			want: 1,
			ok:   true,
		},
		"daisy": {
			want: 2,
			ok:   true,
		},
		"koopa": {
			want: 3,
			ok:   true,
		},
		"lakitu": {
			want: 4,
			ok:   true,
		},
		"luigi": {
			want: 5,
			ok:   true,
		},
		"mario": {
			want: 6,
			ok:   true,
		},
		"peach": {
			want: 7,
			ok:   true,
		},
		"rosalina": {
			want: 8,
			ok:   true,
		},
		"toad": {
			want: 9,
			ok:   true,
		},
		"wario": {
			want: 10,
			ok:   true,
		},
		"yoshi": {
			want: 11,
			ok:   true,
		},
		"lumalee": {
			want: -1,
			ok:   false,
		},
	}
	sort.Strings(values)

	for valueToSearch, data := range tests {
		t.Run(valueToSearch, func(tt *testing.T) {
			givenValue := valueToSearch
			want := data.want
			exist := data.ok
			tt.Parallel()
			// When
			got, ok := slices.BinarySearch(values, givenValue)
			// Then
			if ok != exist {
				tt.Fatalf("%s: want: %d to exist: %t", givenValue, want, exist)
			}

			if !exist {
				return
			}

			if want != got {
				tt.Errorf("%s: want: %d, but got: %d", givenValue, want, got)
			}
		})
	}
}

func TestClip(t *testing.T) {
	t.Parallel()
	// Given
	values := make([]int, 3, 5)
	for i := 1; i <= 3; i++ {
		values[i-1] = i
	}
	wantedSize := 3
	wantedCap := 3
	// When
	values = slices.Clip(values)
	// Then
	if wantedSize != len(values) {
		t.Errorf("want: %d, but got: %d", wantedSize, len(values))
	}
	if wantedCap != cap(values) {
		t.Errorf("want: %d, but got: %d", wantedCap, cap(values))
	}
}

func TestClone(t *testing.T) {
	t.Parallel()
	// Given
	values := make([]int, 3, 5)
	for i := 1; i <= 3; i++ {
		values[i-1] = i
	}
	wantedSize := 3
	wantedCap := 3
	// When
	values = slices.Clip(values)
	// Then
	if wantedSize != len(values) {
		t.Errorf("want: %d, but got: %d", wantedSize, len(values))
	}
	if wantedCap != cap(values) {
		t.Errorf("want: %d, but got: %d", wantedCap, cap(values))
	}
}
