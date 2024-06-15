package http

type Encoder interface {
	Encode(body string) string
}

var Encoding = map[string]Encoder{
	"gzip": nil,
}
