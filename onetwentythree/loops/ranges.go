package loops

import (
	"iter"
)

func DoubleValues(values []int) []int {
	result := make([]int, 0, len(values))
	for v := range buildIterator(values) {
		result = append(result, v)
	}

	return result
}

func buildIterator(values []int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for _, v := range values {
			// fmt.Println(v)
			if !yield(v * 2) {
				return
			}
		}
	}
}

type MyStructs []MyStruct

func (ms MyStructs) Names() iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, s := range ms {
			if !yield(s.Name) {
				return
			}
		}
	}
}

type MyStruct struct {
	ID   int
	Name string
}
