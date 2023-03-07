package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloEndpoint(t *testing.T) {
	/*
		First we'll use go's built in http mocking tools to create
		a mock request and a mock response writer
	*/
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	w := httptest.NewRecorder()

	/*
		Then we'll call the endpoint handler directly with our mocks
	*/
	hello(w, req)
	/*
		now we can read the response from our handler.
		I'm using the 'defer' keyword which will defer the execution
		of that line until the function closes. This way I can
		be sure that the result body will always be closed.

		Go is able to have multiple return types.
		In this case (as in most), it's an error
	*/
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)

	/*
		Now we can check that the results are as we expect
		and fail the test if they are not
	*/
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "hello" {
		t.Errorf("expected hello got %v", string(data))
	}
}
