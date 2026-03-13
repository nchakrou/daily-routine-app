package main

import (
	"daily-routine-backend/internal/db"
	"daily-routine-backend/internal/server"
	"fmt"
	"log"
	"net/http"
)

const (
	DBPath = "./app.db"
	Port   = ":8080"
)

func main() {
	db, err := db.Init(DBPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("server started on http://localhost:8080")
	srv := server.New(db)

	if err := http.ListenAndServe(Port, srv); err != nil {
		log.Fatal(err)
	}
}
