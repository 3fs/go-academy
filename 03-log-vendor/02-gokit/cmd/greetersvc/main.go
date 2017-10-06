package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/3fs/go-academy/03-log-vendor/02-gokit/pkg/greeter"
)

func startServer(address string, handler http.Handler) {
	fmt.Println("Starting server on http://" + address)
	http.ListenAndServe(address, handler)
}

func main() {
	var addr = flag.String("addr", "127.0.0.1:8080", "Interface and port to listen on")

	// parse the flags
	flag.Parse()

	service := greeter.New()

	endpoint := greeter.MakeEndpoint(service)

	handler := greeter.NewHTTPHandler(endpoint)

	startServer(*addr, handler)
}
