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

	// Handle encoding
	httpResponse := http.NewMessage().SetEncoder(httpRequest.GetHeader("Accept-Encoding"))

	if httpRequest.Method == "GET" {
		handleGet(conn, filesDir, httpRequest, httpResponse)
	} else if httpRequest.Method == "POST" {
		handlePost(conn, filesDir, httpRequest, httpResponse)
	}
	httpResponse.Write(conn)
}

func handlePost(conn net.Conn, filesDir string, httpRequest, httpResponse *http.Message) {
	filePrefix := "/files/"
	if strings.HasPrefix(httpRequest.URL, filePrefix) {
		fileName := strings.TrimPrefix(httpRequest.URL, filePrefix)
		fullPath := path.Join(filesDir, fileName)
		err := writeFile(fullPath, httpRequest.Body)
		if err != nil {
			httpResponse.SetStatus(500)
			return
		}
		httpResponse.SetStatus(201)
	} else {
		httpResponse.SetStatus(404)
	}
	return
}

func handleGet(conn net.Conn, filesDir string, httpRequest, httpResponse *http.Message) {

	echoPrefix := "/echo/"
	filePrefix := "/files/"
	if httpRequest.URL == "/" {
		httpResponse.SetStatus(200)
	} else if strings.HasPrefix(httpRequest.URL, echoPrefix) {
		body := strings.TrimPrefix(httpRequest.URL, echoPrefix)
		httpResponse.SetStatus(200).SetHeader("Content-Type", "text/plain").SetBody(body)
	} else if httpRequest.URL == "/user-agent" {
		body := httpRequest.GetHeader("User-Agent")
		httpResponse.SetStatus(200).SetHeader("Content-Type", "text/plain").SetBody(body)
	} else if strings.HasPrefix(httpRequest.URL, filePrefix) {
		fileName := strings.TrimPrefix(httpRequest.URL, filePrefix)
		fullPath := path.Join(filesDir, fileName)
		content, err := readFile(fullPath)
		if err != nil {
			if os.IsNotExist(err) {
				// Handle the case where the file does not exist
				responseMessage := http.NewMessage().SetStatus(404)
				responseMessage.Write(conn)
			}
			return
		}
		httpResponse.SetStatus(200).SetHeader("Content-Type", "application/octet-stream").SetBody(content)

	} else {
		httpResponse.SetStatus(404)
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

func writeFile(filename string, content string) error {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}
