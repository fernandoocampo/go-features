package routes_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/fernandoocampo/go-features/onetwentytwo/routes"
	"github.com/stretchr/testify/require"
)

func TestCreatePet(t *testing.T) {
	// Given
	router := routes.NewRouter()
	webServer := httptest.NewServer(router)
	defer webServer.Close()

	path, err := url.JoinPath(webServer.URL, "/pets")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		t.FailNow()
	}

	want := routes.Pet{
		Name: "drila",
		ID:   "5",
	}
	newPet := routes.Pet{
		Name: "drila",
	}
	requestBody, err := json.Marshal(newPet)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	// When
	got, err := http.Post(path, "application/json", bytes.NewReader(requestBody))
	// Then
	if err != nil {
		t.Errorf("unexpected error calling endpoint: %s", err)
	}

	defer got.Body.Close()

	gotBody, err := io.ReadAll(got.Body)
	if err != nil {
		t.Errorf("unexpected error reading body: %s", err)
	}

	if got.StatusCode > 299 {
		t.Errorf("unexpected response: [%d] %s", got.StatusCode, string(gotBody))
	}

	var savedPet routes.Pet
	err = json.Unmarshal(gotBody, &savedPet)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		t.FailNow()
	}

	if want != savedPet {
		t.Errorf("want: %s, but got: %+v", want, savedPet)
	}
}

func TestGetPet(t *testing.T) {
	// Given
	router := routes.NewRouter()
	want := routes.Pet{
		ID:   "5",
		Name: "Drila",
	}

	webServer := httptest.NewServer(router)
	defer webServer.Close()

	path, err := url.JoinPath(webServer.URL, "/pets")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		t.FailNow()
	}

	// When
	got, err := http.DefaultClient.Get(path)
	// Then
	if err != nil {
		t.Errorf("unexpected error calling endpoint: %s", err)
	}

	defer got.Body.Close()

	gotBody, err := io.ReadAll(got.Body)
	require.NoError(t, err)

	if got.StatusCode > 299 {
		t.Errorf("unexpected response: [%d] %s", got.StatusCode, string(gotBody))
		t.FailNow()
	}

	var result routes.Pet
	err = json.Unmarshal(gotBody, &result)
	require.NoError(t, err)

	if want != result {
		t.Errorf("want: %+v, but got: %+v", want, string(gotBody))
	}
}

func TestGetPetByID(t *testing.T) {
	// Given
	router := routes.NewRouter()
	want := routes.Pet{
		ID:   "5",
		Name: "ddddd",
	}
	webServer := httptest.NewServer(router)
	defer webServer.Close()

	path, err := url.JoinPath(webServer.URL, "/pets/5")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		t.FailNow()
	}

	// When
	got, err := http.Get(path)
	// Then
	if err != nil {
		t.Errorf("unexpected error calling endpoint: %s", err)
	}

	defer got.Body.Close()

	gotBody, err := io.ReadAll(got.Body)
	if err != nil {
		t.Errorf("unexpected error reading body: %s", err)
	}

	var result routes.Pet
	err = json.Unmarshal(gotBody, &result)
	require.NoError(t, err)

	if want != result {
		t.Errorf("want: %+v, but got: %+v", want, result)
	}
}

func TestUpdatePet(t *testing.T) {
	// Given
	router := routes.NewRouter()
	want := routes.Pet{
		ID:   "5",
		Name: "Drila",
	}
	webServer := httptest.NewServer(router)
	path, err := url.JoinPath(webServer.URL, "/pets")
	require.NoError(t, err)

	updatedPet := routes.Pet{
		ID:   "5",
		Name: "Drila",
	}

	updatedPetBytes, err := json.Marshal(updatedPet)
	require.NoError(t, err)

	updatePetRequest, err := http.NewRequest(http.MethodPut, path, bytes.NewReader(updatedPetBytes))
	if err != nil {
		t.Fatalf("creating update pet request: %s", err)
	}
	// When
	got, err := http.DefaultClient.Do(updatePetRequest)
	// Then
	if err != nil {
		t.Errorf("unexpected error calling endpoint: %s", err)
	}

	defer got.Body.Close()

	gotBody, err := io.ReadAll(got.Body)
	if err != nil {
		t.Errorf("unexpected error reading body: %s", err)
	}

	var result routes.Pet
	err = json.Unmarshal(gotBody, &result)
	require.NoError(t, err)

	if want != result {
		t.Errorf("want: %+v, but got: %+v", want, result)
	}
}
