package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_handlers(t *testing.T) {
	var theTests = []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{name: "home", url: "/", expectedStatusCode: http.StatusOK},
		{name: "404", url: "/fish", expectedStatusCode: http.StatusNotFound},
	}

	var app application
	routes := app.routes()

	// create a test server
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	pathToTemplate = "./../../templates/"

	// range through test data
	for _, e := range theTests {
		resp, err := ts.Client().Get(ts.URL + e.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("for %s: expected status %d, but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		}
	}
}
