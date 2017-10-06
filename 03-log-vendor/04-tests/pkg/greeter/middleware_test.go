package greeter

import (
	"bytes"
	"testing"

	"github.com/go-kit/kit/log"
)

func TestServiceValidatorMiddleware(t *testing.T) {
	s := ValidateMiddleware()(New())

	if expect, got := "", s.Greet(""); expect != got {
		t.Errorf("expected '%v', got '%v'", expect, got)
	}

	if expect, got := "Hello there!", s.Greet("there"); expect != got {
		t.Errorf("expected '%v', got '%v'", expect, got)
	}
}

func TestServiceLoggerMiddleare(t *testing.T) {

	buf := new(bytes.Buffer)
	logger := log.NewLogfmtLogger(buf)

	s := ServiceLoggingMiddleware(logger)(New())

	s.Greet("test")

	if expect, got := "method=Greet name=test\n", buf.String(); expect != got {
		t.Errorf("expected '%v', got '%v'", expect, got)
	}
}
