package greeter

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/go-kit/kit/log"
)

type mockService struct {
	prefix string
}

func (s mockService) Greet(name string) string {
	return s.prefix + "-" + name
}

func TestEndpoint(t *testing.T) {
	service := mockService{
		prefix: "sth",
	}
	endpoint := MakeEndpoint(service)

	tests := []struct {
		name        string
		requestName string
		result      string
	}{
		{
			"No name",
			"",
			"sth-",
		},
		{
			"Name set",
			"test",
			"sth-test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := Request{
				Name: tt.requestName,
			}

			resp, err := endpoint(context.Background(), req)
			if err != nil {
				t.Fatalf("expect error to be nil, got %v", err)
			}

			response := resp.(Response)

			if expect, got := tt.result, response.Greet; expect != got {
				t.Errorf("expected '%v', got '%v'", expect, got)
			}
		})
	}
}

func TestEndpointLogging(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := log.NewLogfmtLogger(buf)

	service := mockService{
		prefix: "sth",
	}
	endpoint := EndpointLoggingMiddleware(logger)(MakeEndpoint(service))

	req := Request{}

	_, err := endpoint(context.Background(), req)
	if err != nil {
		t.Fatalf("expect error to be nil, got %v", err)
	}

	if expect, got := "transport_error=null took=", buf.String(); !strings.HasPrefix(got, expect) {
		t.Errorf("expected '%v', got '%v'", expect, got)
	}

}
