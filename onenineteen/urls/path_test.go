package urls_test

import (
	"net/url"
	"testing"
)

// TestJoinPath test url.JoinPath
// JoinPath returns a URL string with the
// provided path elements joined to the
// existing path of base and the resulting
// path cleaned of any ./ or ../ elements.
func TestJoinPath(t *testing.T) {
	// Given
	basePath := "http://app.domain.com"
	path1 := "path1"
	path2 := "path2"
	want := "http://app.domain.com/path1/path2"

	// When
	got, err := url.JoinPath(basePath, path1, path2)

	// Then
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if want != got {
		t.Errorf("want: %s, but got: %s", want, got)
	}
}
