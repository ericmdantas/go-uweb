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
	
}

type UWeb struct {
	node *node
}

func New() *UWeb {
	return &UWeb{}
}

func (uw *UWeb) Get(path string, fn http.HandlerFunc) {
	uw.addNode(path, fn)
}

func (uw *UWeb) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("!"))
}

func (uw *UWeb) addNode(path string, handler http.HandlerFunc) {
	var n map[string]http.HandlerFunc
	uw.node = append(uw.node, handler})
}
