package main

import (
	"net/http"
)

func todo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

func main() {
	http.NewServeMux()
	http.HandleFunc("/", todo)

	http.ListenAndServe(":3000", nil)
}
