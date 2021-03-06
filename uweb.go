package uweb

import "net/http"

type UWebHandlerFunc func(w http.ResponseWriter, r *http.Request)

type UWeb struct {
	nodeMap map[string]Node
}

func New() UWeb {
	return UWeb{
		nodeMap: make(map[string]Node),
	}
}

func (uw UWeb) GET(path string, fn UWebHandlerFunc) {
	uw.addNode(GET, path, fn)
}

func (uw UWeb) POST(path string, fn UWebHandlerFunc) {
	uw.addNode(POST, path, fn)
}

func (uw UWeb) PUT(path string, fn UWebHandlerFunc) {
	uw.addNode(PUT, path, fn)
}

func (uw UWeb) HEAD(path string, fn UWebHandlerFunc) {
	uw.addNode(HEAD, path, fn)
}

func (uw UWeb) PATCH(path string, fn UWebHandlerFunc) {
	uw.addNode(PATCH, path, fn)
}

func (uw UWeb) DELETE(path string, fn UWebHandlerFunc) {
	uw.addNode(DELETE, path, fn)
}

func (uw UWeb) OPTIONS(path string, fn UWebHandlerFunc) {
	uw.addNode(OPTIONS, path, fn)
}

func (uw UWeb) CONNECT(path string, fn UWebHandlerFunc) {
	uw.addNode(CONNECT, path, fn)
}

func (uw UWeb) TRACE(path string, fn UWebHandlerFunc) {
	uw.addNode(TRACE, path, fn)
}

func (uw UWeb) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	n := uw.findHandler(r)

	if n != nil {
		n.handle(w, r)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func (uw UWeb) addNode(method, path string, handlerFn UWebHandlerFunc) {
	pathNormalized := normalizePath(path)
	uw.nodeMap[method+" "+pathNormalized] = newNode(method, pathNormalized, handlerFn)
}

func (uw UWeb) findHandler(r *http.Request) *Node {
	for _, v := range uw.nodeMap {
		if v.isItForMe(r) {
			return &v
		}
	}

	return nil
}
