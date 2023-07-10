package main

import (
	"database/sql"
	"log"
	"net/http"
	"ledgerApp/src/utils/routes"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost/myLedgerApp?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := routes.NewRouter(db)

	log.Fatal(http.ListenAndServe(":8080", router))
}
