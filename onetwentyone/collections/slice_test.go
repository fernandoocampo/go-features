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
	values := []int{0, 1, 2, 3, 4, 5, 6}
	want := []int{0, 1, 2, 3, 4, 5, 6}
	// When
	clonedSlice := slices.Clone(values)
	values[3] = -3
	// Then
	if len(want) != len(clonedSlice) {
		t.Fatalf("want: %v, but got: %v", want, clonedSlice)
	}
	for index := range want {
		if want[index] != clonedSlice[index] {
			t.Errorf("want: %v, but got: %v", want, clonedSlice)
		}
	}
	if values[3] == clonedSlice[3] {
		t.Errorf("original slice at index 3 was changed: %d, but got: %d", values[3], clonedSlice[3])
	}
}

func TestCompact(t *testing.T) {
	t.Parallel()
	// Given
	// Repeated values must be consecutive
	values := []int{0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 6}
	want := []int{0, 1, 2, 3, 4, 5, 6}
	// When
	got := slices.Compact(values)
	// Then
	for index := range want {
		if want[index] != got[index] {
			t.Fatalf("want: %v, but got: %v", want, got)
		}
	}
}

func TestCompareEqual(t *testing.T) {
	t.Parallel()
	// Given
	a := []int{0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 6}
	b := []int{0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 6}
	want := 0
	// When
	got := slices.Compare(a, b)
	// Then
	if got != want {
		t.Errorf("slices are expected to be the same but they are not")
	}
}

func TestCompareDifferent(t *testing.T) {
	t.Parallel()
	// Given
	a := []int{0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 6}
	b := []int{0, 1, 1, 2, 2, 3, 3, 4, 4, 5}
	awant := 1
	bwant := -1
	// When
	agot := slices.Compare(a, b)
	bgot := slices.Compare(b, a)
	// Then
	if agot != awant {
		t.Errorf("want: %d, but got: %d", awant, agot)
	}
	if bgot != bwant {
		t.Errorf("want: %d, but got: %d", bwant, bgot)
	}
}

func TestContains(t *testing.T) {
	t.Parallel()
	// Given
	a := []int{0, 1, 2, 3, 4, 5}
	b := 3
	want := true
	// When
	got := slices.Contains(a, b)
	// Then
	if want != got {
		t.Errorf("want: %v, but got: %v", want, got)
	}
}

func TestDelete(t *testing.T) {
	t.Parallel()
	// Given
	a := []int{0, 1, 2, 3, 4, 5}
	want := []int{0, 1, 4, 5}
	// When
	// [i:j]
	got := slices.Delete(a, 2, 4)
	// Then
	if slices.Compare(got, want) != 0 {
		t.Errorf("want: %v, but got: %v", want, got)
	}
}

func TestEqual(t *testing.T) {
	t.Parallel()
	// Given
	a := []int{0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 6}
	b := []int{0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 6}
	want := true
	// When
	got := slices.Equal(a, b)
	// Then
	if got != want {
		t.Errorf("slices are expected to be equal but they are not")
	}
}

func TestEqualDifferentCapacity(t *testing.T) {
	t.Parallel()
	// Given
	a := make([]int, 10, 20)
	b := make([]int, 10, 21)
	want := true
	// When
	got := slices.Equal(a, b)
	// Then
	if got != want {
		t.Errorf("slices are expected to be equal but they are not")
	}
}

func TestGrow(t *testing.T) {
	t.Parallel()
	// Given
	values := make([]int, 0, 5)
	want := make([]int, 0, 20)
	// When
	values = slices.Grow(values, 20)
	// Then
	if len(values) != len(want) {
		t.Errorf("want length: %d, but got: %d", len(want), len(values))
	}
	if cap(values) != cap(want) {
		t.Errorf("want capacity: %d, but got: %d", cap(want), cap(values))
	}
}

func TestIndex(t *testing.T) {
	t.Parallel()
	// Given
	values := []rune{0: 'q', 23: 'r', 34: 'z'}
	want := 23
	// When
	got := slices.Index(values, 'r')
	// Then
	if want != got {
		t.Errorf("want: %d, but got: %d", want, got)
	}
}

func TestInsert(t *testing.T) {
	t.Parallel()
	// Given
	atIndex := 5
	newValue := 23
	values := []int{0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 6}
	want := []int{0, 1, 1, 2, 2, 23, 3, 3, 4, 4, 5, 6}
	// When
	got := slices.Insert(values, atIndex, newValue)
	// Then
	if !slices.Equal(want, got) {
		t.Errorf("want: %v, but got: %v", want, got)
	}
}

func TestSorted(t *testing.T) {
	t.Parallel()
	// Given
	values := []int{0, 1, 2, 3, 4, 5, 6, 7, 99}
	want := true
	// When
	got := slices.IsSorted(values)
	// Then
	if want != got {
		t.Errorf("want: %t, but got: %t", want, got)
	}
}

func TestNotSorted(t *testing.T) {
	t.Parallel()
	// Given
	values := []int{0, 1, 2, 3, 44, 5, 6, 7, 99}
	want := false
	// When
	got := slices.IsSorted(values)
	// Then
	if want != got {
		t.Errorf("want: %t, but got: %t", want, got)
	}
}

func TestMax(t *testing.T) {
	t.Parallel()
	// Given
	values := []int{0, 1, 2, 3, 44, 5, 103, 6, 7, 99}
	want := 103
	// When
	got := slices.Max(values)
	// Then
	if want != got {
		t.Errorf("want: %d, but got: %d", want, got)
	}
}

func TestMin(t *testing.T) {
	t.Parallel()
	// Given
	values := []int{0, 1, 2, 3, -1, 44, 5, 103, 6, 7, 99}
	want := -1
	// When
	got := slices.Min(values)
	// Then
	if want != got {
		t.Errorf("want: %d, but got: %d", want, got)
	}
}

func TestMinAllEquals(t *testing.T) {
	t.Parallel()
	// Given
	values := make([]int, 20)
	want := 0
	// When
	got := slices.Min(values)
	// Then
	if want != got {
		t.Errorf("want: %d, but got: %d", want, got)
	}
}

func TestReplace(t *testing.T) {
	t.Parallel()
	// Given
	fromIndex := 2
	toIndex := 7
	newValues := []int{2, 3, 4, 5, 6}
	values := []int{0, 1, 1, 2, 2, 3, 3, 7, 8, 9, 10}
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// When
	got := slices.Replace(values, fromIndex, toIndex, newValues...)
	// Then
	if !slices.Equal(want, got) {
		t.Errorf("want: %v, but got: %v", want, got)
	}
}

func TestReverse(t *testing.T) {
	t.Parallel()
	// Given
	values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	want := []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	// When
	slices.Reverse(values)
	// Then
	if !slices.Equal(want, values) {
		t.Errorf("want: %v, but got: %v", want, values)
	}
}

func TestSort(t *testing.T) {
	t.Parallel()
	// Given
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	values := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	// When
	slices.Sort(values)
	// Then
	if !slices.Equal(want, values) {
		t.Errorf("want: %v, but got: %v", want, values)
	}
}
