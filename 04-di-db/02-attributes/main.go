package main

import (
	"fmt"
	"net/http"

	"github.com/3fs/go-academy/04-di-db/01-initial/db"
	"github.com/3fs/go-academy/04-di-db/01-initial/log"
)

func main() {
	myDB, _ := db.New("prod.db.com")
	logger, _ := log.New()

	http.HandleFunc("/", rootHandler(myDB, logger))
	http.ListenAndServe(":8080", nil)
}

func rootHandler(myDB *db.DB, logger *log.Log) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Log("rootHandler invoked")
		result, _ := myDB.Read("rootElement")
		fmt.Fprintf(w, "Found %s", result)
	}
}
