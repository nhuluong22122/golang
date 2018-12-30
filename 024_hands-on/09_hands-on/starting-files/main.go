package main

import (
	"log"
	"net/http"
	"text/template"
)

//# Serve the files in the "starting-files" folder
//
//## To get your images to serve, use:
//
//``` Go
//	func StripPrefix(prefix string, h Handler) Handler
//	func FileServer(root FileSystem) Handler
//```
//
//Constraint: you are not allowed to change the route being used for images in the template file
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	//fs := http.FileServer(http.Dir("public"))
	http.HandleFunc("/", dogs)
	http.Handle("/public/",  http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	//http.HandleFunc("/", dogs)
	http.ListenAndServe(":8080", nil)
}

func dogs(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
