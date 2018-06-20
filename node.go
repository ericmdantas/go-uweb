package uweb

import (
	"net/http"
	"strings"
)

const (
	maskIdentifier = ":"
)

type Node struct {
	path   string
	method string
	handle UWebHandlerFunc
}

func (n Node) isItForMe(r *http.Request) bool {
	if r == nil {
		return false
	}

	if n.method != r.Method {
		return false
	}

	if n.path == r.URL.Path {
		return true
	}

	if !strings.Contains(n.path, maskIdentifier) {
		return false
	}

	nodePathPartsWithtoutSlashes := strings.Split(n.path, "/")
	requestPathPartsWithtoutSlashes := strings.Split(r.URL.Path, "/")
	var mountedPathParts []string
	var mountedPathString string

	for index, nodePathPart := range nodePathPartsWithtoutSlashes {
		if strings.Contains(nodePathPart, maskIdentifier) {
			mountedPathParts = append(mountedPathParts, requestPathPartsWithtoutSlashes[index])
		} else {
			mountedPathParts = append(mountedPathParts, nodePathPartsWithtoutSlashes[index])
		}
	}

	mountedPathString = strings.Join(mountedPathParts, "/")

	return mountedPathString == r.URL.Path
}

func newNode(method, path string, fn UWebHandlerFunc) *Node {
	return &Node{
		method: normalizeMethod(method),
		path:   normalizePath(path),
		handle: fn,
	}
}
