package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Starting HTTP server in port 4221!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	HandleRequest(conn)
}

func HandleRequest(conn net.Conn) {
	httpMessage := HTTPMessage{}
	httpMessage.Unmarshal(conn)

	fmt.Printf("%+v\n", httpMessage)

	echoPrefix := "/echo/"
	if httpMessage.Status.URL == "/" {
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else if strings.HasPrefix(httpMessage.Status.URL, echoPrefix) {
		text := strings.TrimPrefix(httpMessage.Status.URL, echoPrefix)
		response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %v\r\n\r\n%v", len(text), text)
		conn.Write([]byte(response))
	} else if httpMessage.Status.URL == "/user-agent" {
		text := httpMessage.Headers["User-Agent"]
		response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %v\r\n\r\n%v", len(text), text)
		conn.Write([]byte(response))
	} else {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}
}
