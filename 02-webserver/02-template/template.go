package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `<form method="post">
		<input name="name" value="{{ .Name }}" required> <input type="submit" value="Greet!">
		</form>
		{{ if .Name }}<h1>Hello {{ .Name }}!</h1>{{ end }}`

	// set the encoding
	w.Header().Add("Content-type", "text/html")

	// validate the method
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// prepare the data
	data := struct {
		Name string
	}{
		Name: r.FormValue("name"),
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

func startServer() {
	http.HandleFunc("/", htmlHandler)

	fmt.Println("Starting server on http://0.0.0.0:8080")
	http.ListenAndServe(":8080", nil)
}

func main() {
	startServer()
}
