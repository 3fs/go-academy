package greeter

import (
	"context"
	"strings"
	"time"

	"github.com/go-kit/kit/endpoint"
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
	if strings.TrimSpace(name) == "" {
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

// EndpointLoggingMiddleware returns an endpoint middleware that logs the
// duration of each invocation, and the resulting error, if any.
func EndpointLoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			defer func(begin time.Time) {
				logger.Log("transport_error", err, "took", time.Since(begin))
			}(time.Now())
			return next(ctx, request)

		}
	}
}
