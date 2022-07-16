package main

import (
	"net/http"
	"text/template"
)

type Todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Items []Todo
}

var tmpl := *template.


func todo(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "TODO LIST",
		Items: []Todo{
			{Item: "Instal GO", Done: true},
			{Item: "Learn GO", Done: false},
			{Item: "Teach GO", Done: true},
		},
	}
}

func main() {
	s := http.NewServeMux()
	s.HandleFunc("/hello", todo)
	
	http.ListenAndServe(":3000", s)

}
