package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"

	"github.com/3fs/go-academy/02-webserver/07-makefile/pkg/greeter"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	// validate method
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// read and validate the query string
	name := r.FormValue("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// greet the user
	fmt.Fprint(w, greeter.Greet(name))
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `<form method="post">
		<input name="name" required> <input type="submit" value="Greet!">
		</form>
		{{ if .Greet }}<h1>{{ .Greet }}</h1>{{ end }}`

	// set the encoding
	w.Header().Add("Content-type", "text/html")

	// validate the method
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	greet := ""
	if r.FormValue("name") != "" {
		greet = greeter.Greet(r.FormValue("name"))
	}

	// prepare the data
	data := struct {
		Greet string
	}{
		Greet: greet,
	}

	// parse the template
	t, err := template.New("form").Parse(tmpl)
	if err != nil {
		fmt.Println("Failed to parse template;", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t.Execute(w, data)
}

func startServer(address string) {
	http.HandleFunc("/api/greet", apiHandler)
	http.HandleFunc("/greet", htmlHandler)

	fmt.Println("Starting server on http://" + address)
	http.ListenAndServe(address, nil)
}

func main() {
	var addr = flag.String("addr", "", "Interface and port to listen on")

	// parse the flags
	flag.Parse()

	startServer(*addr)
}
