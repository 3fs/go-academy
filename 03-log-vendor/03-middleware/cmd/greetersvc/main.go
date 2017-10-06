package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"

	"github.com/3fs/go-academy/03-log-vendor/03-middleware/pkg/greeter"
)

func main() {
	var addr = flag.String("addr", "127.0.0.1:8080", "Interface and port to listen on")

	// parse the flags
	flag.Parse()

	// Create a single logger, which we'll use and give to other components.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// Create greeter service
	var service greeter.Service
	{
		service = greeter.New()
		service = greeter.ServiceLoggingMiddleware(log.With(logger, "service", "greeter"))(service)
		service = greeter.ValidateMiddleware()(service)
	}

	endpoint := greeter.MakeEndpoint(service)

	handler := greeter.NewHTTPHandler(endpoint, logger)

	logger.Log("transport", "http", "listen", *addr)
	err := http.ListenAndServe(*addr, handler)
	if err != nil {
		logger.Log("transport", "http", "during", "listen", "err", err)
	}
}
