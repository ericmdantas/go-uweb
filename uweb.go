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

type node struct {
	method, path string
	handlerFn    http.HandlerFunc
}

type UWeb struct {
	tree []node
}

func New() *UWeb {
	return &UWeb{}
}

func (uw *UWeb) Get(path string, fn http.HandlerFunc) {
	uw.addNode(get, path, fn)
}

func (uw *UWeb) Post(path string, fn http.HandlerFunc) {
	uw.addNode(post, path, fn)
}

func (uw *UWeb) Put(path string, fn http.HandlerFunc) {
	uw.addNode(put, path, fn)
}

func (uw *UWeb) Patch(path string, fn http.HandlerFunc) {
	uw.addNode(patch, path, fn)
}

func (uw *UWeb) Delete(path string, fn http.HandlerFunc) {
	uw.addNode(delete, path, fn)
}

func (uw *UWeb) Options(path string, fn http.HandlerFunc) {
	uw.addNode(options, path, fn)
}

func (uw *UWeb) Connect(path string, fn http.HandlerFunc) {
	uw.addNode(connect, path, fn)
}

func (uw *UWeb) Trace(path string, fn http.HandlerFunc) {
	uw.addNode(trace, path, fn)
}

func (uw *UWeb) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	n := uw.findHandler(r)
	n.handlerFn(w, r)
}

func (uw *UWeb) addNode(method, path string, handlerFn http.HandlerFunc) {
	uw.tree = append(uw.tree, node{method: method, path: path, handlerFn: handlerFn})
}

func (uw *UWeb) findHandler(r *http.Request) node {
	for _, v := range uw.tree {
		if (v.path == r.URL.Path) && (v.method == r.Method) {
			return v
		}
	}

	return node{} // should be nil
}
