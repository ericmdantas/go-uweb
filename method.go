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

func normalizeMethod(verb string) string {
	v := strings.TrimSpace(strings.ToUpper(verb))

	if len(v) == 0 {
		panic("No method specified")
	}

	if (v != GET) && (v != HEAD) && (v != POST) &&
		(v != PUT) && (v != PATCH) && (v != DELETE) &&
		(v != CONNECT) && (v != OPTIONS) && (v != TRACE) {
		panic("Invalid verb: " + v)
	}

	return v
}
