package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})

	err := http.ListenAndServeTLS(":3000", ".cer/server.crt", ".cer/server.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}
