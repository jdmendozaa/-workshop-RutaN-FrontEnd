package http

import (
	"bytes"
	"compress/gzip"
)

type Encoder interface {
	Encode(body string) (string, error)
}

var Encoding = map[string]Encoder{
	"gzip": Gzip{},
}

type Gzip struct{}

func (Gzip) Encode(body string) (string, error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	_, err := gz.Write([]byte(body))
	if err != nil {
		return "", err
	}
	if err := gz.Close(); err != nil {
		return "", err
	}
	return b.String(), nil
}
