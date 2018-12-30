package main

import (
	"net/http"
)

//# Serve the files in the "starting-files" folder
//
//Use "http.FileServer"
func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./starting-files"))))
	//log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
	http.ListenAndServe(":8080", nil)
}
