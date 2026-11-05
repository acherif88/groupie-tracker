package api

import "testing"

func TestBaseURLParDefaut(t *testing.T) {
	t.Setenv("API_URL", "")
	if got := baseURL(); got != defaultBaseURL {
		t.Errorf("attendu %s, obtenu %s", defaultBaseURL, got)
	}
}

func TestBaseURLDepuisVariable(t *testing.T) {
	t.Setenv("API_URL", "http://localhost:9000/api/")
	if got := baseURL(); got != "http://localhost:9000/api" {
		t.Errorf("attendu http://localhost:9000/api, obtenu %s", got)
	}
}
