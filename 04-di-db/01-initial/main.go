package main

import (
	"fmt"
	"net/http"

	"github.com/3fs/go-academy/04-di-db/01-initial/db"
	"github.com/3fs/go-academy/04-di-db/01-initial/log"
)

var myDB *db.DB
var logger *log.Log

func main() {
	myDB, _ = db.New("prod.db.com")
	logger, _ = log.New()

	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log("rootHandler invoked")
	result, _ := myDB.Read("rootElement")
	fmt.Fprintf(w, "Found %s", result)
}
