package routes_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fernandoocampo/go-features/onetwentytwo/routes"
)

func TestCreatePet(t *testing.T) {
	// Given
	router := routes.NewRouter()
	want := "post a pet"
	webServer := httptest.NewServer(router)
	path := fmt.Sprintf("%s/pets", webServer.URL)
	// When
	got, err := http.Post(path, "text/plain", nil)
	// Then
	if err != nil {
		t.Errorf("unexpected error calling endpoint: %s", err)
	}

	gotBody, err := io.ReadAll(got.Body)
	if err != nil {
		t.Errorf("unexpected error reading body: %s", err)
	}

	defer got.Body.Close()

	if want != string(gotBody) {
		t.Errorf("want: %s, but got: %s", want, string(gotBody))
	}
}

func TestGetPet(t *testing.T) {
	// Given
	router := routes.NewRouter()
	want := "get pets"
	webServer := httptest.NewServer(router)
	path := fmt.Sprintf("%s/pets", webServer.URL)
	// When
	got, err := http.Get(path)
	// Then
	if err != nil {
		t.Errorf("unexpected error calling endpoint: %s", err)
	}

	gotBody, err := io.ReadAll(got.Body)
	if err != nil {
		t.Errorf("unexpected error reading body: %s", err)
	}

	defer got.Body.Close()

	if want != string(gotBody) {
		t.Errorf("want: %s, but got: %s", want, string(gotBody))
	}
}

func TestGetPetByID(t *testing.T) {
	// Given
	router := routes.NewRouter()
	want := "get pet: 1"
	webServer := httptest.NewServer(router)
	path := fmt.Sprintf("%s/pets/1", webServer.URL)
	// When
	got, err := http.Get(path)
	// Then
	if err != nil {
		t.Errorf("unexpected error calling endpoint: %s", err)
	}

	gotBody, err := io.ReadAll(got.Body)
	if err != nil {
		t.Errorf("unexpected error reading body: %s", err)
	}

	defer got.Body.Close()

	if want != string(gotBody) {
		t.Errorf("want: %s, but got: %s", want, string(gotBody))
	}
}

func TestUpdatePet(t *testing.T) {
	// Given
	router := routes.NewRouter()
	want := "put a pet"
	webServer := httptest.NewServer(router)
	path := fmt.Sprintf("%s/pets", webServer.URL)
	updatePetRequest, err := http.NewRequest(http.MethodPut, path, nil)
	if err != nil {
		t.Fatalf("creating update pet request: %s", err)
	}
	// When
	got, err := http.DefaultClient.Do(updatePetRequest)
	// Then
	if err != nil {
		t.Errorf("unexpected error calling endpoint: %s", err)
	}

	gotBody, err := io.ReadAll(got.Body)
	if err != nil {
		t.Errorf("unexpected error reading body: %s", err)
	}

	defer got.Body.Close()

	if want != string(gotBody) {
		t.Errorf("want: %s, but got: %s", want, string(gotBody))
	}
}
