package uweb

import "net/http"

const (
	get     = "GET"
	head    = "HEAD"
	post    = "POST"
	put     = "PUT"
	patch   = "PATCH"
	delete  = "DELETE"
	connect = "CONNECT"
	options = "OPTIONS"
	trace   = "TRACE"
)

type UWebHandlerFunc func(w http.ResponseWriter, r *http.Request)

type UWeb struct {
	tree []*Node
}

func New() *UWeb {
	return &UWeb{}
}

func (uw *UWeb) Get(path string, fn UWebHandlerFunc) {
	uw.addNode(get, path, fn)
}

func (uw *UWeb) Post(path string, fn UWebHandlerFunc) {
	uw.addNode(post, path, fn)
}

func (uw *UWeb) Put(path string, fn UWebHandlerFunc) {
	uw.addNode(put, path, fn)
}

func (uw *UWeb) Patch(path string, fn UWebHandlerFunc) {
	uw.addNode(patch, path, fn)
}

func (uw *UWeb) Delete(path string, fn UWebHandlerFunc) {
	uw.addNode(delete, path, fn)
}

func (uw *UWeb) Options(path string, fn UWebHandlerFunc) {
	uw.addNode(options, path, fn)
}

func (uw *UWeb) Connect(path string, fn UWebHandlerFunc) {
	uw.addNode(connect, path, fn)
}

func (uw *UWeb) Trace(path string, fn UWebHandlerFunc) {
	uw.addNode(trace, path, fn)
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
