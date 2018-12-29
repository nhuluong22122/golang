package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

//# Create a basic server using TCP.
//
//The server should use net.Listen to listen on port 8080.
//
//Remember to close the listener using defer.
//
//Remember that from the "net" package you first need to LISTEN, then you need to ACCEPT an incoming connection.
//
//Now write a response back on the connection.
//
//Use io.WriteString to write the response: I see you connected.
//
//Remember to close the connection.
//
//Once you have all of that working, run your TCP server and test it from telnet (telnet localhost 8080).
func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
		log.Println(err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn){
	fmt.Println("CONNECTED")
	scanner := bufio.NewScanner(conn)
	//use this to read the connection
	for scanner.Scan() { // on the default give you the line of text
		ln := scanner.Text()
		fmt.Println(ln)
	}

	defer conn.Close()
	//The reason why we never get to here because we were caught in that scanner.Scan() loop that always try to read in new inputs
	//and never close
	fmt.Println("TEST")
	io.WriteString(conn, "I see you connected.")

	conn.Close()

}
