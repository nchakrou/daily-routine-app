package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("server started on http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
