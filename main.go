package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := os.Getenv("NAME")

		fmt.Fprintf(w, "Hello World!, testing docker")
		fmt.Fprintf(w, "Nama : ", name)
	})

	http.ListenAndServe(":8080", nil)
}
