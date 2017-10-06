package greeter

import (
	"github.com/go-kit/kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(Service) Service

// ValidateMiddleware is a validator middleware for service
func ValidateMiddleware() Middleware {
	return func(next Service) Service {
		return validateMiddleware{next}
	}
}

type validateMiddleware struct {
	next Service
}

func (mw validateMiddleware) Greet(name string) string {
	if name == "" {
		return ""
	}

	return mw.next.Greet(name)
}

// ServiceLoggingMiddleware is a validator middleware for service
func ServiceLoggingMiddleware(log log.Logger) Middleware {
	return func(next Service) Service {
		return loggingMiddleware{log, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

func (mw loggingMiddleware) Greet(name string) string {
	mw.logger.Log("method", "Greet", "name", name)
	return mw.next.Greet(name)
}
