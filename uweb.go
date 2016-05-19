package uweb

import "net/http"

type UWebHandlerFunc func(w http.ResponseWriter, r *http.Request)

type UWeb struct {
	tree []*Node
}

func New() *UWeb {
	return &UWeb{}
}

func (uw *UWeb) Get(path string, fn UWebHandlerFunc) {
	uw.addNode(GET, path, fn)
}

func (uw *UWeb) Post(path string, fn UWebHandlerFunc) {
	uw.addNode(POST, path, fn)
}

func (uw *UWeb) Put(path string, fn UWebHandlerFunc) {
	uw.addNode(PUT, path, fn)
}

func (uw *UWeb) Patch(path string, fn UWebHandlerFunc) {
	uw.addNode(PATCH, path, fn)
}

func (uw *UWeb) Delete(path string, fn UWebHandlerFunc) {
	uw.addNode(DELETE, path, fn)
}

func (uw *UWeb) Options(path string, fn UWebHandlerFunc) {
	uw.addNode(OPTIONS, path, fn)
}

func (uw *UWeb) Connect(path string, fn UWebHandlerFunc) {
	uw.addNode(CONNECT, path, fn)
}

func (uw *UWeb) Trace(path string, fn UWebHandlerFunc) {
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
