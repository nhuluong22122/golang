package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)
func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
		log.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			log.Println(err)
			continue
		}
		go serve(conn)
	}

}

func serve(conn net.Conn){
	fmt.Println("CONNECTED")
	scanner := bufio.NewScanner(conn)
	//use this to read the connection
	var i int
	var rMethod string
	var rURI string
	for scanner.Scan() { // on the default give you the line of text
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// we're in REQUEST LINE
			xs := strings.Fields(ln)
			rMethod = xs[0]
			rURI = xs[1]
			fmt.Println("METHOD:", rMethod)
			fmt.Println("URI:", rURI)
		}
		if len(ln) == 0 { break } //to break out of the loop -> TEST gets to print
		i++
 	}
	body := "<h1>HOLY COW THIS IS LOW LEVEL</h1>"
	body += "<br>"
	body += rMethod
	body += "<br>"
	body += rURI
	body += "<br>"
	io.WriteString(conn,"HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn,"\r\n")
	io.WriteString(conn, body)

	conn.Close()

}
