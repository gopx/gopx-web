package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("New connection from: %s\n", r.RemoteAddr)
		fmt.Fprintf(w, "Hi, I'm GoPX")
	})

	log.Println("Server is listening on 8080 port...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
