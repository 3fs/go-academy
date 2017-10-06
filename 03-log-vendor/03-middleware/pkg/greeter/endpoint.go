package greeter

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Request definition
type Request struct {
	Name string
}

// Response definition
type Response struct {
	Greet string
}

// MakeEndpoint creates endpoint for greeter
func MakeEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(Request)

		return Response{
			Greet: s.Greet(req.Name),
		}, nil
	}
}
