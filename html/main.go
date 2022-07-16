package main

import (
	"fmt"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func hi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1 color=red>Hello</h1>"))
}
func GetBook(w http.ResponseWriter, r *http.Request) {
	books := Book{
		Title:  "Hello from json",
		Author: "John",
		Pages:  1,
	}
	w.Write([]byte(fmt.Sprintf("Title:%v\nAuthor:%v", books.Title, books.Author)))
}

func main() {
	http.HandleFunc("/", hi)
	http.HandleFunc("/Hello", Hello)
	http.HandleFunc("/book", GetBook)
	http.ListenAndServe(":3000", nil)
}
