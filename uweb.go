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

func (uw *UWeb) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, v := range uw.tree {
		if v.path == r.URL.Path {
			v.handlerFn(w, r)
		}
	}
}

func (uw *UWeb) addNode(method, path string, handlerFn http.HandlerFunc) {
	uw.tree = append(uw.tree, node{method: method, path: path, handlerFn: handlerFn})
}
