package newerrors_test

import (
	"errors"
	"testing"

	"github.com/fernandoocampo/go-features/onetwenty/newerrors"
)

func TestMultipleErrors(t *testing.T) {
	t.Parallel()
	// Given
	testcases := map[string]struct {
		anUser       newerrors.User
		wantedErrors []error
	}{
		"all": {
			anUser: newerrors.User{},
			wantedErrors: []error{
				newerrors.ErrInvalidFirstName,
				newerrors.ErrInvalidLastName,
				newerrors.ErrInvalidUserName,
				newerrors.ErrInvalidEmail,
			},
		},
		"username": {
			anUser: newerrors.User{
				FirstName: "michael",
				LastName:  "Jordan",
				Email:     "michael.jordan@sports.com",
			},
			wantedErrors: []error{
				newerrors.ErrInvalidUserName,
			},
		},
		"email": {
			anUser: newerrors.User{
				FirstName: "michael",
				LastName:  "Jordan",
				Username:  "michael.jordan",
			},
			wantedErrors: []error{
				newerrors.ErrInvalidEmail,
			},
		},
		"first_name": {
			anUser: newerrors.User{
				LastName: "Jordan",
				Username: "michael.jordan",
				Email:    "michael.jordan@sports.com",
			},
			wantedErrors: []error{
				newerrors.ErrInvalidFirstName,
			},
		},
		"last_name": {
			anUser: newerrors.User{
				FirstName: "michael",
				Username:  "michael.jordan",
				Email:     "michael.jordan@sports.com",
			},
			wantedErrors: []error{
				newerrors.ErrInvalidLastName,
			},
		},
	}
	// When
	for title, value := range testcases {
		testName := title
		testData := value

		t.Run(title, func(tt *testing.T) {
			got := testData.anUser.Validate()

			// Then
			if got == nil && len(testData.wantedErrors) > 0 {
				tt.Errorf("%q: errors were expected but got nil", testName)
			}

			for _, wantedError := range testData.wantedErrors {
				if !errors.Is(got, wantedError) {
					tt.Errorf("%q: want %q, but got: %s", testName, wantedError, got)
				}
			}
		})
	}
}
