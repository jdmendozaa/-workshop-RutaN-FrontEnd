package http

var StatusCode = map[int]string{
	200: "OK",
	404: "Not Found",
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	500: "Internal Server Error",
	502: "Bad Gateway",
	503: "Service Unavailable",
}
