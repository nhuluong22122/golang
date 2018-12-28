package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)
//create a multiplexer (mux, servemux, router, server, http mux, ...)
// so that your server responds to different URIs and METHODS.

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	request(conn)

}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			m := strings.Fields(ln)[0]
			u := strings.Fields(ln)[1]
			t := strings.Fields(ln)[2]

			fmt.Println("***METHOD", m)
			fmt.Println("***URL", u)
			fmt.Println("***HTTP TYPE", t)
			respond(conn, m, u)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}


func respond(conn net.Conn, method string, url string) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>` + conn.LocalAddr().String() + ` ` + method + ` ` + url + `</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
