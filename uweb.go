package uweb

import (
	"net/http"
)

const (
	GET     = "GET"
	HEAD    = "HEAD"
	POST    = "POST"
	PUT     = "PUT"
	PATCH   = "PATCH"
	DELETE  = "DELETE"
	CONNECT = "CONNECT"
	OPTIONS = "OPTIONS"
	TRACE   = "TRACE"
)

type UWeb struct {
}

func New() *UWeb {
	return &UWeb{}
}

func (uw *UWeb) Get(path string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}
}

func (uw *UWeb) Post(path string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}
}

func (uw *UWeb) Put(path string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}
}

func (uw *UWeb) Delete(path string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}
}

func (uw *UWeb) Patch(path string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}
}

func (uw *UWeb) Head(path string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}
}

func (uw *UWeb) Options(path string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}
}

func (uw *UWeb) Run(port string) error {
	return http.ListenAndServe(port, nil)
}
