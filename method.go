package uweb

import (
	"strings"
)

const (
	GET     = "GET"
	HEAD    = "HEAD"
	POST    = "POST"
	PUT     = "PUT"
	PATCH   = "PATCH"
	DELETE  = "DELETE"
	CONNECT = "CONNECT"
	OPTIONS = "OPTIONS"
	TRACE   = "TRACE"
)

func NormalizeMethod(verb string) string {
	v := strings.ToUpper(verb)
	v = strings.TrimSpace(v)

	if length := len(v); length == 0 || length > 7 {
		panic("Invalid verb: " + v)
	}

	if (v != GET) && (v != HEAD) && (v != POST) && (v != PUT) && (v != PATCH) && (v != DELETE) && (v != CONNECT) && (v != OPTIONS) && (v != TRACE) {
		panic("Invalid verb: " + v)
	}

	return v
}
