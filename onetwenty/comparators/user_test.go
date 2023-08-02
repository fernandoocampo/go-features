package comparators_test

import (
	"testing"

	"github.com/fernandoocampo/go-features/onetwenty/comparators"
)

func TestCompareUsers(t *testing.T) {
	t.Parallel()
	// Given
	givenUser := comparators.User{
		Name:     "Gitlberto",
		UserName: "gitlbert.osorio",
		Age:      23,
	}

	testcases := map[string]struct {
		user comparators.User
		want bool
	}{
		"false": {
			user: comparators.User{
				Name:     "Gitlbertos",
				UserName: "gitlbert.osorio",
				Age:      23,
			},
			want: false,
		},
		"true": {
			user: comparators.User{
				Name:     "Gitlberto",
				UserName: "gitlbert.osorio",
				Age:      23,
			},
			want: true,
		},
	}
	// When
	for title, value := range testcases {
		testData := value
		t.Run(title, func(tt *testing.T) {
			tt.Parallel()
			got := comparators.Equal(testData.user, givenUser)
			// Then
			if got != testData.want {
				tt.Errorf("user: %+v is not equal to %+v", testData.user, givenUser)
			}
		})
	}
}
