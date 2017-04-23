package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alinex/go-learn/cmd"
)

func Test_Status(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	cmd.Status(res, req)

	exp := "OK"
	act := res.Body.String()
	if exp != act {
		t.Fatalf("Expected %s got %s", exp, act)
	}
}
