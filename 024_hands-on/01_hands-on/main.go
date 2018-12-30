package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

//# ListenAndServe on port 8080 of localhost
//
//For the default route "/"
//Have a func called "foo"
//which writes to the response "foo ran"
//
//For the route "/dog/"
//Have a func called "dog"
//which parses a template called "dog.gohtml"
//and writes to the response "<h1>This is from dog</h1>"
//and also shows a picture of a dog when the template is executed.
//
//Use "http.ServeFile"
//to serve the file "dog.jpeg"

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/dog.jpg", dogFile) //Always require a different api for pictures

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(w, nil);
}
func dogFile(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog.jpg")
}
