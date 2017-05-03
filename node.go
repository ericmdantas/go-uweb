package uweb

type Node struct {
	path   string
	method string
	handle UWebHandlerFunc
}

func newNode(method, path string, fn UWebHandlerFunc) *Node {
	return &Node{
		method: normalizeMethod(method),
		path:   normalizePath(path),
		handle: fn,
	}
}
