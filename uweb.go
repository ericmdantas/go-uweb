package uweb

import "net/http"

type UWebHandlerFunc func(w http.ResponseWriter, r *http.Request)

type UWeb struct {
	tree []*Node
}

func New() *UWeb {
	return &UWeb{}
}

func (uw *UWeb) GET(path string, fn UWebHandlerFunc) {
	uw.addNode(GET, path, fn)
}

func (uw *UWeb) POST(path string, fn UWebHandlerFunc) {
	uw.addNode(POST, path, fn)
}

func (uw *UWeb) PUT(path string, fn UWebHandlerFunc) {
	uw.addNode(PUT, path, fn)
}

func (uw *UWeb) HEAD(path string, fn UWebHandlerFunc) {
	uw.addNode(HEAD, path, fn)
}

func (uw *UWeb) PATCH(path string, fn UWebHandlerFunc) {
	uw.addNode(PATCH, path, fn)
}

func (uw *UWeb) DELETE(path string, fn UWebHandlerFunc) {
	uw.addNode(DELETE, path, fn)
}

func (uw *UWeb) OPTIONS(path string, fn UWebHandlerFunc) {
	uw.addNode(OPTIONS, path, fn)
}

func (uw *UWeb) CONNECT(path string, fn UWebHandlerFunc) {
	uw.addNode(CONNECT, path, fn)
}

func (uw *UWeb) TRACE(path string, fn UWebHandlerFunc) {
	uw.addNode(TRACE, path, fn)
}

func (uw *UWeb) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	n := uw.findHandler(r)
	n.handle(w, r)
}

func (uw *UWeb) addNode(method, path string, handlerFn UWebHandlerFunc) {
	n := newNode(method, path, handlerFn)
	uw.tree = append(uw.tree, n)
}

func (uw *UWeb) findHandler(r *http.Request) *Node {
	for _, n := range uw.tree {
		if (n.path == r.URL.Path) && (n.method == r.Method) {
			return n
		}
	}

	return nil
}
