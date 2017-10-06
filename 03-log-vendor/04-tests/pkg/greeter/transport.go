package greeter

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func decodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	name := r.FormValue("name")

	return Request{
		Name: name,
	}, nil
}

// writes response from endpoint to client
func encodePlainResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(Response)

	w.Header().Add("Content-type", "text/plain")

	_, err := fmt.Fprint(w, resp.Greet)
	return err
}

// writes decorated response from endpoint to client
func encodeHTMLResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(Response)

	tmpl := `<form method="post">
		<input name="name" required> <input type="submit" value="Greet!">
		</form>
		{{ if .Greet }}<h1>{{ .Greet }}</h1>{{ end }}`

	w.Header().Add("Content-type", "text/html")

	// prepare the data
	data := struct {
		Greet string
	}{
		Greet: resp.Greet,
	}

	t, err := template.New("form").Parse(tmpl)
	if err != nil {
		return err
	}

	return t.Execute(w, data)
}

// NewHTTPHandler creates greeter handlers
func NewHTTPHandler(endpoint endpoint.Endpoint, logger log.Logger) http.Handler {
	m := http.NewServeMux()

	m.Handle("/api/greet", httptransport.NewServer(
		EndpointLoggingMiddleware(log.With(logger, "endpoint", "api"))(endpoint),
		decodeRequest,
		encodePlainResponse,
	))

	m.Handle("/greet", httptransport.NewServer(
		EndpointLoggingMiddleware(log.With(logger, "endpoint", "html"))(endpoint),
		decodeRequest,
		encodeHTMLResponse,
	))

	return m
}
