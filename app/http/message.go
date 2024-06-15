package http

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)

type Message struct {
	Method      string
	URL         string
	HTTPVersion string
	StatusCode  int
	StatusText  string
	Headers     map[string]string
	Body        string
	Encoder     Encoder
}

func Unmarshal(conn net.Conn) (*Message, error) {

	reader := bufio.NewReader(conn)
	statusString, _ := reader.ReadString('\n')
	statusArr := strings.Fields(statusString)

	headers := make(map[string]string)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		if strings.TrimSpace(line) == "" {
			break
		}
		header := strings.Split(line, ":")
		// We create a map with each element of the header
		headers[strings.TrimSpace(header[0])] = strings.TrimSpace(header[1])
	}
	body := ""
	if bodyLength, ok := headers["Content-Length"]; ok {
		bodyLengthI, err := strconv.Atoi(bodyLength)
		if err != nil {
			return nil, err
		}
		bodyB, _ := reader.Peek(bodyLengthI)
		body = string(bodyB)
	}

	httpMessage := Message{
		Method:      statusArr[0],
		URL:         statusArr[1],
		HTTPVersion: statusArr[2],
		Headers:     headers,
		Body:        body,
	}

	return &httpMessage, nil
}

func (httpMessage *Message) Marshal() ([]byte, error) {
	status := fmt.Sprintf("%v %v %v", httpMessage.HTTPVersion, httpMessage.StatusCode, httpMessage.StatusText)
	// Create headers string
	var headers string
	for key, value := range httpMessage.Headers {
		headers += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	response := fmt.Sprintf("%v\r\n%v\r\n%v", status, headers, httpMessage.Body)
	return []byte(response), nil
}

func NewMessage() *Message {
	httpMessage := Message{
		HTTPVersion: "HTTP/1.1",
		Headers:     make(map[string]string),
	}
	httpMessage.SetHeader("Content-Length", "0")
	return &httpMessage
}

func (httpMessage *Message) SetHeader(key, value string) *Message {
	httpMessage.Headers[key] = value
	return httpMessage
}

func (httpMessage *Message) GetHeader(key string) string {
	return httpMessage.Headers[key]
}

func (httpMessage *Message) SetBody(body string) *Message {
	if httpMessage.Encoder != nil {
		// TODO: Handle errors
		encodedBody, _ := httpMessage.Encoder.Encode(body)
		httpMessage.Body = encodedBody
	} else {
		fmt.Println("Not Encoding body...")
		httpMessage.Body = body
	}
	httpMessage.SetHeader("Content-Length", fmt.Sprintf("%v", len(httpMessage.Body)))
	return httpMessage
}

func (httpMessage *Message) SetStatus(status int) *Message {
	httpMessage.StatusCode = status
	httpMessage.StatusText = StatusCode[status]
	return httpMessage
}

func (httpMessage *Message) SetEncoder(encoders string) *Message {
	encodersSlice := strings.Split(encoders, ",")
	for _, encoder := range encodersSlice {
		encoder = strings.TrimSpace(encoder)
		if value, ok := Encoding[encoder]; ok {
			httpMessage.Encoder = value
			httpMessage.SetHeader("Content-Encoding", encoder)
			break
		}
	}
	// In case there's already an uncompressed body, we set it again to compress it.
	httpMessage.SetBody(httpMessage.Body)
	return httpMessage
}

func (httpMessage *Message) Write(conn io.Writer) {
	message, _ := httpMessage.Marshal()
	conn.Write(message)
}
