package main

import (
	"html/template"
	"net/http"
)

// Write a simple web server that can be used to test the functionality of your

type server struct {
}

// Start the server
func (s server) start() {

	fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()
	// Create a new http server
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", s.handle)
	// Start the server
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		println(err.Error())
	}
}

func (s server) handle(writer http.ResponseWriter, _ *http.Request) {
	// Write a response
	tpl := template.Must(template.ParseFiles("assets/index.html"))
	err := tpl.Execute(writer, nil)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	// Create a new server
	s := &server{}
	// Start the server
	s.start()
}
