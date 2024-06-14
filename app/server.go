package main

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/http"
	"net"
	"os"
	"path"
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
	filePath := ""
	if len(os.Args) > 2 {
		filePath = os.Args[2]
	}

	fmt.Println("FilePath:", filePath)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go HandleRequest(conn, filePath)
	}
}

func HandleRequest(conn net.Conn, filesDir string) {
	httpRequest, _ := http.Unmarshal(conn)

	echoPrefix := "/echo/"
	filePrefix := "/files/"
	if httpRequest.URL == "/" {
		responseMessage := http.NewMessage().SetStatus(200)
		message, _ := responseMessage.Marshal()
		conn.Write(message)
	} else if strings.HasPrefix(httpRequest.URL, echoPrefix) {
		body := strings.TrimPrefix(httpRequest.URL, echoPrefix)
		responseMessage := http.NewMessage().SetStatus(200).SetHeader("Content-Type", "text/plain").SetBody(body)
		message, _ := responseMessage.Marshal()
		conn.Write(message)
	} else if httpRequest.URL == "/user-agent" {
		body := httpRequest.Headers["User-Agent"]
		responseMessage := http.NewMessage().SetStatus(200).SetHeader("Content-Type", "text/plain").SetBody(body)
		message, _ := responseMessage.Marshal()
		conn.Write(message)
	} else if strings.HasPrefix(httpRequest.URL, filePrefix) {
		fileName := strings.TrimPrefix(httpRequest.URL, filePrefix)
		fullPath := path.Join(filesDir, fileName)
		content, err := readFile(fullPath)
		if err != nil {
			if os.IsNotExist(err) {
				// Handle the case where the file does not exist
				responseMessage := http.NewMessage().SetStatus(404)
				message, _ := responseMessage.Marshal()
				conn.Write(message)
			}
			return
		}
		responseMessage := http.NewMessage().SetStatus(200).SetHeader("Content-Type", "application/octet-stream").SetBody(content)
		message, _ := responseMessage.Marshal()
		conn.Write(message)

	} else {
		responseMessage := http.NewMessage().SetStatus(404)
		message, _ := responseMessage.Marshal()
		conn.Write(message)
	}
}

func readFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	fmt.Println(len(content))
	if err != nil {
		return "", err
	}
	return string(content), nil
}
