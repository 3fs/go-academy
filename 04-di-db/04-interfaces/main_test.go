package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockDb struct {
	t            *testing.T
	input        string
	outputString string
	outputError  error
}

func (m *mockDb) Read(input string) (string, error) {
	// validate input
	if input != m.input {
		t.Errorf("Expect read with '%s', got '%s'", m.input, input)
		return "", fmt.Errorf("Expect read with '%s', got '%s'", m.input, input)
	}

	return m.outputString, m.outputError
}

type mockLogger struct{}

func (m *mockLogger) Log(string) {}

func TestRootHandler(t *testing.T) {
	// setup
	handlers := handlers{
		db:  &mockDb{t, "rootElement", "root", nil},
		log: &mockLogger{},
	}
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	// call it
	handlers.root(rr, req)

	// check it
	if rr.Body.String() != "Found root" {
		t.Errorf("Expected response \"%s\" to equal \"Found root\"", rr.Body.String())
	}
}

func TestMain(m *testing.M) {
	m.Run()
}
