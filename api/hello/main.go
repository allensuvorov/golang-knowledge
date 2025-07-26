package main

import (
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello from the cloud - start")
	w.Write([]byte("Hello from the cloud!"))
	log.Println("Hello from the cloud - end")
}

func main() {
	// Register the two new handler functions and corresponding route patterns with
	// the servemux, in exactly the same way that we did before.
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
