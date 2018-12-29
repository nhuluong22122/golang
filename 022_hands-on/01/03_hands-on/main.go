package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
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

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}


func dog(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "index.gohtml", "dog dog dog")
	HandleError(res, err)
}

func mainPage(res http.ResponseWriter, req *http.Request){
	err := tpl.ExecuteTemplate(res, "index.gohtml", "main page")
	HandleError(res, err)
}

func me(res http.ResponseWriter, req *http.Request){
	fmt.Println(req.RequestURI)
	name := strings.Split(req.RequestURI, "/");
	err := tpl.ExecuteTemplate(res, "index.gohtml", name[2] + " PAGE")
	HandleError(res, err)
	//io.WriteString(res,"USER " + name[2])
}

func main() {
	http.HandleFunc("/", mainPage) //can use HandleFunc
	http.HandleFunc("/dog", dog) //can use HandleFunc
	http.HandleFunc("/me/", me) //can use HandleFunc

	http.ListenAndServe(":8080", nil)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
