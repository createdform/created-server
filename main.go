package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var PORT = 80

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Service is healthy"))
	})

	http.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadFile("commit_hash.txt")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("unable to read file: " + err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	})

	fmt.Printf("Server is running on port %v\n", PORT)
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
	fmt.Printf("Server stopped: %v\n", err)
}
