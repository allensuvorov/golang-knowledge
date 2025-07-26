package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"time"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	resp := HelloResponse{Message: "Hello, Lightning AI!"}
	json.NewEncoder(w).Encode(resp)
}

func stream(w http.ResponseWriter, r *http.Request) {
		// Make sure the ResponseWriter supports flushing
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	// Set headers to prevent buffering in proxies or curl
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Cache-Control", "no-cache")

	// Write + Flush in chunks
	w.Write([]byte("Part 1\n"))
	flusher.Flush()
	time.Sleep(1 * time.Second)

	w.Write([]byte("Part 2\n"))
	flusher.Flush()
	time.Sleep(1 * time.Second)

	w.Write([]byte("Part 3\n"))
	flusher.Flush()
	time.Sleep(1 * time.Second)
}

// {
//   "name": "Alice"
// }

func greet(w http.ResponseWriter, r *http.Request) {
	var v struct {
		Name string
	}

	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if v.Name == "" {
		http.Error(w, "Missing 'name', field", http.StatusBadRequest)
	}

	obj := HelloResponse {
		Message: fmt.Sprintf("Hello, %s!", v.Name),
	}

	json.NewEncoder(w).Encode(obj)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", hello)
	mux.HandleFunc("GET /stream", stream)
	mux.HandleFunc("POST /greet", greet)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
