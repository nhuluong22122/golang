package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", read)
	//http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {

}

func read(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.SetCookie(w, &http.Cookie{
			Name:  "my-cookie",
			Value: "1",
			Path: "/",
		})
		fmt.Fprintln(w, "NUMBER OF TIME VISITED A SITE:", 1)
		return
	}
	update, err := strconv.Atoi(c.Value)
	update = update + 1
	c.Value = strconv.Itoa(update)
	http.SetCookie(w, c)
	fmt.Fprintln(w, "NUMBER OF TIME VISITED A SITE:", c.Value)

}

// Using cookies, track how many times a user has been to your website domain.