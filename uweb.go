package uweb

import "net/http"

type UWebHandlerFunc func(w http.ResponseWriter, r *http.Request)

type UWeb struct {
	tree []Node
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

	if n != nil {
		n.handle(w, r)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404"))
}

func (uw *UWeb) addNode(method, path string, handlerFn UWebHandlerFunc) {
	uw.tree = append(uw.tree, newNode(method, path, handlerFn))
}

func (uw *UWeb) findHandler(r *http.Request) *Node {
	for _, n := range uw.tree {
		if n.isItForMe(r) {
			return &n
		}
	}

	return nil
}
