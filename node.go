package uweb

import "net/http"

type Node struct {
	path   string
	method string
	handle UWebHandlerFunc
}

func (n Node) isItForMe(r *http.Request) bool {
	if r == nil {
		return false
	}

	return n.path == r.URL.Path && n.method == r.Method
}

func newNode(method, path string, fn UWebHandlerFunc) *Node {
	return &Node{
		method: normalizeMethod(method),
		path:   normalizePath(path),
		handle: fn,
	}
}
