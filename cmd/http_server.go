package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerfunction)
	http.ListenAndServe(":1722", nil)
}

func handlerfunction(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("Hello %s", r.URL.Path[1:])))
}
