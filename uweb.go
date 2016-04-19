package uweb

import (
	"fmt"
	"net/http"
)

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

type UWeb struct {
}

func New() *UWeb {
	return &UWeb{}
}

func (uw *UWeb) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	w.Write([]byte("!"))
}
