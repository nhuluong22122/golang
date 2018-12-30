package main

import (
	"net/http"
)

//# Serve the files in the "starting-files" folder
//
//## To get your images to serve, use only this:
//
//``` Go
//	fs := http.FileServer(http.Dir("public"))
//```
//
//Hint: look to see what type FileServer returns, then think it through.

func main() {
	//http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./starting-files"))))
	fs :=  http.FileServer(http.Dir("public"))
	http.Handle("/pics/", fs)
	http.ListenAndServe(":8080", nil)
}