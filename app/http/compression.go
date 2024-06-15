package http

import (
	"bytes"
	"compress/gzip"
)

type Compressor interface {
	Compress(body string) (string, error)
}

var Algorithms = map[string]Compressor{
	"gzip": Gzip{},
}

type Gzip struct{}

func (Gzip) Compress(body string) (string, error) {
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
