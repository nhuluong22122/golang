package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
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
	for scanner.Scan() { // on the default give you the line of text
		ln := scanner.Text()
		fmt.Println(ln)
		if len(ln) == 0 { break } //to break out of the loop -> TEST gets to print
 	}

	fmt.Println("TEST")
	fmt.Fprintln(conn,"test WRITE to connection using fmt")
	io.WriteString(conn, "test WRITE to connection using io.")

	conn.Close()

}
