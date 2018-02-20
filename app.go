package main

import (
	"fmt"
	"net/http"
	"os"
)

func handleRequest(w http.ResponseWriter, req *http.Request) {
	host, err := os.Hostname()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf("Sending this response from %s \n", host)))
}

func main() {
	http.HandleFunc(`/`, handleRequest)
	http.ListenAndServe(`:8080`, nil)
}
