package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func __TestHandler(t *testing.T) {
	testText := "Bienvenido, testing"
	http.HandleFunc("/", handler)
	go func() {
		err := http.ListenAndServe(":80", nil)
		if err != nil {
			t.Error()
		}
	}()
	resp, err := http.Get("http://127.0.0.1:80/?key=testing")
	if err != nil {
		t.Error()
	}
	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error()
	}
	if string(responseData) != testText {
		t.Error()
	}
}

func TestHealthCheckHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/?key=testing", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.ç
	expected := "Bienvenido, testing"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
