package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type HTTPStatus struct {
	Method      string
	URL         string
	HTTPVersion string
}

type HTTPMessage struct {
	Status  HTTPStatus
	Headers map[string]string
	Body    string
}

func (httpMessage *HTTPMessage) Unmarshal(conn net.Conn) {

	reader := bufio.NewReader(conn)
	statusString, _ := reader.ReadString('\n')
	statusArr := strings.Fields(statusString)
	httpStatus := HTTPStatus{
		Method:      statusArr[0],
		URL:         statusArr[1],
		HTTPVersion: statusArr[2],
	}
	headers := make(map[string]string)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(line) == "" {
			break
		}
		header := strings.Split(line, ":")
		// We create a map with each element of the header
		headers[strings.TrimSpace(header[0])] = strings.TrimSpace(header[1])
	}

	*httpMessage = HTTPMessage{
		Status:  httpStatus,
		Headers: headers,
		// For now, I don't care about the body
		Body: "",
	}
}
