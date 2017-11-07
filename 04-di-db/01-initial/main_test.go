package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/3fs/go-academy/04-di-db/01-initial/db"
	"github.com/3fs/go-academy/04-di-db/01-initial/log"
)

func TestRootHandler(t *testing.T) {
	myDB, _ = db.New("test.db.com")
	logger, _ = log.NewDummyLog()

	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	// call it
	rootHandler(rr, req)

	if rr.Body.String() != "Found root" {
		t.Errorf("Expected response \"%s\" to equal \"Found root\"", rr.Body.String())
	}
}

func TestMain(m *testing.M) {
	m.Run()
}
