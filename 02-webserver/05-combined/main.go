package main

import (
	"bufio"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type apiHandler struct{}

func (*apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	fmt.Fprintf(w, "Hello %s!\n", name)
}

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

func startServer(address string) {
	http.Handle("/api/greet", &apiHandler{})
	http.HandleFunc("/greet", htmlHandler)

	fmt.Println("Starting server on http://" + address)
	http.ListenAndServe(address, nil)
}

func readStdin() string {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return ""
	}

	if fi.Mode()&os.ModeNamedPipe == 0 {
		return ""
	}

	s, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		return ""
	}

	return string(s)
}

func main() {
	var (
		addr = flag.String("addr", "", "Interface and port to listen on")
		name = flag.String("name", "", "Name of the person you'd like to greet")
	)

	// parse the flags
	flag.Parse()

	if *addr != "" {
		startServer(*addr)
	} else {
		greet := readStdin()
		if greet == "" {
			greet = *name
		}

		fmt.Printf("Hello %s!\n", greet)
	}
}
