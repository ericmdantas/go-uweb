package uweb

import (
	"net/url"
)

func newNode(method, path string, fn UWebHandlerFunc) *Node {
	return &Node{
		method: method,
		path:   NormalizePath(path),
		handle: fn,
	}
}

type Node struct {
	path        string
	method      string
	parsed      *url.URL
	pathParams  map[string]string
	queryString map[string][]string
	handle      UWebHandlerFunc
}

func (n *Node) Parse() {
	n.parsed, _ = url.Parse(n.path)
}

func (n *Node) Query() map[string][]string {
	return n.parsed.Query()
}

func (n *Node) PathParam(k string) map[string]string {
	var q map[string]string

	return q
}
