package uweb

import (
	"net/http"
	"net/url"
	"testing"
)

var tableNewNode = []struct {
	path, wantpath, method, wantmethod string
	wantfnhandler, fnhandler           UWebHandlerFunc
}{
	{
		path:       "/alo",
		wantpath:   "/alo",
		method:     "GET",
		wantmethod: "GET",
		fnhandler:  func(w http.ResponseWriter, r *http.Request) {},
	},
	{
		path:       "/alo1",
		wantpath:   "/alo1",
		method:     "POST",
		wantmethod: "POST",
		fnhandler:  func(w http.ResponseWriter, r *http.Request) {},
	},
}

var tableIsItForMe = []struct {
	method    string
	path      string
	fnhandler UWebHandlerFunc
	req       *http.Request
	want      bool
}{
	{
		method:    "GET",
		path:      "/alo1",
		fnhandler: func(w http.ResponseWriter, r *http.Request) {},
		req:       nil,
		want:      false,
	},
	{
		method:    "GET",
		path:      "/alo1",
		fnhandler: func(w http.ResponseWriter, r *http.Request) {},
		req: &http.Request{
			Method: "GET",
			URL: &url.URL{
				Path: "/alo2",
			},
		},
		want: false,
	},
	{
		method:    "GET",
		path:      "/alo1",
		fnhandler: func(w http.ResponseWriter, r *http.Request) {},
		req: &http.Request{
			Method: "HEAD",
			URL: &url.URL{
				Path: "/alo1",
			},
		},
		want: false,
	},
	{
		method:    "GET",
		path:      "/alo1",
		fnhandler: func(w http.ResponseWriter, r *http.Request) {},
		req: &http.Request{
			Method: "GET",
			URL: &url.URL{
				Path: "/alo1",
			},
		},
		want: true,
	},
	{
		method:    "GET",
		path:      "/aaaaa/:wat",
		fnhandler: func(w http.ResponseWriter, r *http.Request) {},
		req: &http.Request{
			Method: "GET",
			URL: &url.URL{
				Path: "/yeah/aaaaa",
			},
		},
		want: false,
	},
	{
		method:    "GET",
		path:      "/aaaaa/:wat.json",
		fnhandler: func(w http.ResponseWriter, r *http.Request) {},
		req: &http.Request{
			Method: "GET",
			URL: &url.URL{
				Path: "/yeah/aaaaa",
			},
		},
		want: false,
	},
	{
		method:    "GET",
		path:      "/:wat",
		fnhandler: func(w http.ResponseWriter, r *http.Request) {},
		req: &http.Request{
			Method: "GET",
			URL: &url.URL{
				Path: "/yeah",
			},
		},
		want: true,
	},
	{
		method:    "GET",
		path:      "/alo1/:name",
		fnhandler: func(w http.ResponseWriter, r *http.Request) {},
		req: &http.Request{
			Method: "GET",
			URL: &url.URL{
				Path: "/alo1/fulano",
			},
		},
		want: true,
	},
	{
		method:    "GET",
		path:      "/alo1/:name/xyz999",
		fnhandler: func(w http.ResponseWriter, r *http.Request) {},
		req: &http.Request{
			Method: "GET",
			URL: &url.URL{
				Path: "/alo1/fulano/xyz999",
			},
		},
		want: true,
	},
	{
		method:    "GET",
		path:      "/alo1/:name/xyz999/:ident",
		fnhandler: func(w http.ResponseWriter, r *http.Request) {},
		req: &http.Request{
			Method: "GET",
			URL: &url.URL{
				Path: "/alo1/fulano/xyz999/123something",
			},
		},
		want: true,
	},
}

func BenchmarkNewNode(b *testing.B) {
	b.Run("simple", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			newNode("Get", "/", func(w http.ResponseWriter, r *http.Request) {})
		}
	})
}

func TestNewNode(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		for _, v := range tableNewNode {
			n := newNode(v.method, v.path, v.fnhandler)

			if n.path != v.wantpath {
				t.Errorf("Expected path %s, but got: %s", v.wantpath, n.path)
			}

			if n.method != v.wantmethod {
				t.Errorf("Expected method %s, but got: %s", v.wantmethod, n.method)
			}
		}
	})
}

func TestIsItForMe(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		for _, v := range tableIsItForMe {
			n := newNode(v.method, v.path, v.fnhandler)

			r := n.isItForMe(v.req)

			if v.want != r {
				t.Errorf("Request should be resolved %t, but got %t.\nMethod: node %s, req %s.\nPaths: node %s, req %s.", v.want, r, n.method, v.req.Method, n.path, v.req.URL.Path)
			}
		}
	})
}

func BenchmarkIsItForMe(b *testing.B) {
	b.Run("nil_req", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			n := newNode("GET", "/alo", func(w http.ResponseWriter, r *http.Request) {})
			n.isItForMe(nil)
		}
	})

	b.Run("different_method_req", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			n := newNode("GET", "/alo", func(w http.ResponseWriter, r *http.Request) {})
			n.isItForMe(&http.Request{
				Method: "POST",
				URL: &url.URL{
					Path: "/alo",
				},
			})
		}
	})

	b.Run("different_path_req", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			n := newNode("GET", "/alo0", func(w http.ResponseWriter, r *http.Request) {})
			n.isItForMe(&http.Request{
				Method: "GET",
				URL: &url.URL{
					Path: "/alo1",
				},
			})
		}
	})

	b.Run("simple_req", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			n := newNode("GET", "/alo", func(w http.ResponseWriter, r *http.Request) {})
			n.isItForMe(&http.Request{
				Method: "GET",
				URL: &url.URL{
					Path: "/alo",
				},
			})
		}
	})

	b.Run("simple_req_with_querystrings", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			n := newNode("GET", "/alo?a=1", func(w http.ResponseWriter, r *http.Request) {})
			n.isItForMe(&http.Request{
				Method: "GET",
				URL: &url.URL{
					Path: "/alo?a=1",
				},
			})
		}
	})

	b.Run("complex_req", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			n := newNode("GET", "/alo/:name/something/:id/:some_func/:something_else/:wat/:yes", func(w http.ResponseWriter, r *http.Request) {})
			n.isItForMe(&http.Request{
				Method: "GET",
				URL: &url.URL{
					Path: "/alo/a/something/1/fn/123/abc/true",
				},
			})
		}
	})

	b.Run("complex_req_with_querystrings", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			n := newNode("GET", "/alo/:name/something/:id/:some_func/:something_else/:wat/:yes?a=1&b=2&c=3", func(w http.ResponseWriter, r *http.Request) {})
			n.isItForMe(&http.Request{
				Method: "GET",
				URL: &url.URL{
					Path: "/alo/a/something/1/fn/123/abc/true?a=1&b=2&c=3",
				},
			})
		}
	})
}
