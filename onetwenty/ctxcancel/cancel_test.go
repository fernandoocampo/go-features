package ctxcancel_test

import (
	"context"
	"errors"
	"testing"

	"github.com/fernandoocampo/go-features/onetwenty/ctxcancel"
)

func TestCancel(t *testing.T) {
	// Given
	x := 1
	y := 2
	want := errors.New("cause I want")
	ctx, cancel := context.WithCancelCause(context.TODO())
	cause := errors.New("cause I want")
	go func() {
		cancel(cause)
	}()

	// When
	err := ctxcancel.Process(ctx, x, y)

	// Then
	if err == nil {
		t.Fatal("an error was expected but got nil")
	}
	if err.Error() != want.Error() {
		t.Errorf("want: %s, but got: %s", want, err)
	}
}
