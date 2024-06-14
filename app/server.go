package main

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/http"
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
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go HandleRequest(conn)
	}
}

func HandleRequest(conn net.Conn) {
	httpRequest, _ := http.Unmarshal(conn)

	echoPrefix := "/echo/"
	if httpRequest.URL == "/" {
		responseMessage := http.Message{
			HTTPVersion: "HTTP/1.1",
			StatusCode:  200,
			StatusText:  "OK",
		}
		message, _ := responseMessage.Marshal()
		conn.Write(message)
	} else if strings.HasPrefix(httpRequest.URL, echoPrefix) {
		body := strings.TrimPrefix(httpRequest.URL, echoPrefix)
		headers := map[string]string{
			"Content-Type":   "text/plain",
			"Content-Length": fmt.Sprintf("%v", len(body)),
		}
		responseMessage := http.Message{
			HTTPVersion: "HTTP/1.1",
			StatusCode:  200,
			StatusText:  "OK",
			Headers:     headers,
			Body:        body,
		}
		message, _ := responseMessage.Marshal()
		conn.Write(message)
	} else if httpRequest.URL == "/user-agent" {
		body := httpRequest.Headers["User-Agent"]
		headers := map[string]string{
			"Content-Type":   "text/plain",
			"Content-Length": fmt.Sprintf("%v", len(body)),
		}
		responseMessage := http.Message{
			HTTPVersion: "HTTP/1.1",
			StatusCode:  200,
			StatusText:  "OK",
			Headers:     headers,
			Body:        body,
		}
		message, _ := responseMessage.Marshal()
		conn.Write(message)
	} else {
		responseMessage := http.Message{
			HTTPVersion: "HTTP/1.1",
			StatusCode:  404,
			StatusText:  "Not Found",
		}
		message, _ := responseMessage.Marshal()
		conn.Write(message)
	}
}
