package loops_test

import (
	"slices"
	"testing"
	"time"

	"github.com/fernandoocampo/go-features/onetwentytwo/loops"
)

func TestDynamicCounter(t *testing.T) {
	// Given
	countTo := 10
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// When
	got := loops.CountTo(countTo)

	// Then
	if !slices.Equal(want, got) {
		t.Errorf("want: %+v, but got: %+v", want, got)
	}
}

// go test -race -timeout 3s -run ^TestCountLettersWithConcurrency$ github.com/fernandoocampo/go-features/onetwentytwo/loops
func TestCountLettersWithConcurrency(t *testing.T) {
	// Given
	items := []string{
		"Alpha", "Bravo", "Charlie", "Delta", "Echo",
		"Foxtrot", "Golf", "Hotel", "India", "Juliett",
		"Kilo", "Lima", "Mike", "November", "Oscar",
		"Papa", "Quebec", "Romeo", "Sierra", "Tango",
		"Uniform", "Victor", "Whiskey", "X-ray",
		"Yankee", "Zulu",
	}

	done := make(chan struct{})
	want := 140

	go func() {
		<-time.After(1 * time.Second)
		close(done)
	}()

	// When
	got := loops.CountLettersWithConcurrency(done, items)

	// Then
	if want != got {
		t.Errorf("want: %d, but got: %d", want, got)
	}
}

func TestCountLetters(t *testing.T) {
	// Given
	items := []string{
		"Alpha", "Bravo", "Charlie", "Delta", "Echo",
		"Foxtrot", "Golf", "Hotel", "India", "Juliett",
		"Kilo", "Lima", "Mike", "November", "Oscar",
		"Papa", "Quebec", "Romeo", "Sierra", "Tango",
		"Uniform", "Victor", "Whiskey", "X-ray",
		"Yankee", "Zulu",
	}

	want := 140

	// When
	got := loops.CountLetters(items)

	// Then
	if want != got {
		t.Errorf("want: %d, but got: %d", want, got)
	}
}

// "AlphaBravoCharlieDeltaEchoFoxtrotGolfHotelIndiaJuliettKiloLimaMikeNovemberOscarPapaQuebecRomeoSierraTangoUniformVictorWhiskeyX-rayYankeeZulu"
