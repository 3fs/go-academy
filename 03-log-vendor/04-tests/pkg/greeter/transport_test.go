package greeter

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-kit/kit/log"
)

func TestTransport(t *testing.T) {
	endpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Request)
		return Response{
			Greet: req.Name,
		}, nil
	}
	handler := NewHTTPHandler(endpoint, log.NewNopLogger())

	tests := []struct {
		name             string
		path             string
		wantStatusCode   int
		wantBodyContains string
	}{
		{
			"Not found",
			"/not-found",
			http.StatusNotFound,
			"404 page not found",
		},
		{
			"API",
			"/api/greet?name=name-to-test",
			http.StatusOK,
			"name-to-test",
		},
		{
			"HTML",
			"/greet?name=name-to-test",
			http.StatusOK,
			"<h1>name-to-test</h1>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rw := httptest.NewRecorder()

			req, err := http.NewRequest(http.MethodGet, tt.path, nil)

			if err != nil {
				t.Fatalf("expected error to be nil, got '%v'", err)
			}

			// handle request
			handler.ServeHTTP(rw, req)

			if expect, got := tt.wantStatusCode, rw.Code; expect != got {
				t.Errorf("expected '%v', got '%v'", expect, got)
			}

			if expect, got := tt.wantBodyContains, rw.Body.String(); !strings.Contains(got, expect) {
				t.Errorf("expected '%v', got '%v'", expect, got)
			}
		})
	}
}
func TestTransportHTTPClient(t *testing.T) {
	endpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Request)
		return Response{
			Greet: req.Name,
		}, nil
	}

	handler := NewHTTPHandler(endpoint, log.NewNopLogger())

	tests := []struct {
		name             string
		path             string
		wantStatusCode   int
		wantBodyContains string
	}{
		{
			"Not found",
			"/not-found",
			http.StatusNotFound,
			"404 page not found",
		},
		{
			"API",
			"/api/greet?name=name-to-test",
			http.StatusOK,
			"name-to-test",
		},
		{
			"HTML",
			"/greet?name=name-to-test",
			http.StatusOK,
			"<h1>name-to-test</h1>",
		},
	}

	ts := httptest.NewServer(handler)
	defer ts.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req, err := http.NewRequest(http.MethodGet, ts.URL+tt.path, nil)

			if err != nil {
				t.Fatalf("expected error to be nil, got '%v'", err)
			}

			resp, err := http.DefaultClient.Do(req)

			if err != nil {
				t.Fatalf("expected error to be nil, got '%v'", err)
			}

			if expect, got := tt.wantStatusCode, resp.StatusCode; expect != got {
				t.Errorf("expected '%v', got '%v'", expect, got)
			}

			body, _ := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()

			if expect, got := tt.wantBodyContains, string(body); !strings.Contains(got, expect) {
				t.Errorf("expected '%v', got '%v'", expect, got)
			}
		})
	}
}
