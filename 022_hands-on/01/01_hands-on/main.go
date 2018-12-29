package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

//ListenAndServe on port ":8080" using the default ServeMux.
//
//Use HandleFunc to add the following routes to the default ServeMux:
//
//"/"
//"/dog/"
//"/me/
//
//Add a func for each of the routes.
//
//Have the "/me/" route print out your name.


func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func mainPage(res http.ResponseWriter, req *http.Request){
	io.WriteString(res, "main page")
}

func me(res http.ResponseWriter, req *http.Request){
	fmt.Println(req.RequestURI)
	name := strings.Split(req.RequestURI, "/");
	io.WriteString(res,"USER " + name[2])
}

func main() {
	http.HandleFunc("/", mainPage) //can use HandleFunc
	http.HandleFunc("/dog", dog) //can use HandleFunc
	http.HandleFunc("/me/", me) //can use HandleFunc

	http.ListenAndServe(":8080", nil)
}
