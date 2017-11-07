package main

import (
	"fmt"
	"net/http"

	"github.com/3fs/go-academy/04-di-db/01-initial/db"
	"github.com/3fs/go-academy/04-di-db/01-initial/log"
)

type handlers struct {
	db  *db.DB
	log *log.Log
}

func (h *handlers) root(w http.ResponseWriter, r *http.Request) {
	h.log.Log("rootHandler invoked")
	result, _ := h.db.Read("rootElement")
	fmt.Fprintf(w, "Found %s", result)
}

func main() {
	myDB, _ := db.New("prod.db.com")
	logger, _ := log.New()
	h := &handlers{myDB, logger}

	http.HandleFunc("/", h.root)
	http.ListenAndServe(":8080", nil)
}
