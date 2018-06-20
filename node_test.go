package uweb

import (
	"net/http"
	"net/url"
	"testing"
)

var tableNormalizedPath = []struct {
	in, out string
}{
	{
		in:  "",
		out: "/",
	},
	{
		in:  "a",
		out: "/a",
	},
	{
		in:  "B",
		out: "/b",
	},
	{
		in:  "c/",
		out: "/c",
	},
	{
		in:  "D/",
		out: "/d",
	},
	{
		in:  "/E/",
		out: "/e",
	},
	{
		in:  "Abcdef/",
		out: "/abcdef",
	},
	{
		in:  "Abcdef/:id/:s/:x",
		out: "/abcdef/:id/:s/:x",
	},
	{
		in:  "Abcdef/:id/:s/:x",
		out: "/abcdef/:id/:s/:x",
	},
	{
		in:  "/Abcdef/",
		out: "/abcdef",
	},
	{
		in:  "/Abcdef/:id/",
		out: "/abcdef/:id",
	},
	{
		in:  "/Abcdef/:id/          ",
		out: "/abcdef/:id",
	},
	{
		in:  "        /Abcdef/:id/          ",
		out: "/abcdef/:id",
	},
	{
		in:  " /Abcdef/:id/ ",
		out: "/abcdef/:id",
	},
}

var tableNewNodeNotEmptyExistingMethodNormalize = []struct {
	in, out string
}{
	{
		in:  "get",
		out: "GET",
	},
	{
		in:  "head",
		out: "HEAD",
	},
	{
		in:  "post",
		out: "POST",
	},
	{
		in:  "put",
		out: "PUT",
	},
	{
		in:  "options",
		out: "OPTIONS",
	},
	{
		in:  "delete",
		out: "DELETE",
	},
	{
		in:  "connect",
		out: "CONNECT",
	},
	{
		in:  "trace",
		out: "TRACE",
	},
	{
		in:  "patch",
		out: "PATCH",
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
}

func BenchmarkNewNode(b *testing.B) {
	b.Run("simple", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			newNode("Get", "/", func(w http.ResponseWriter, r *http.Request) {})
		}
	})
}

func TestNewNode(t *testing.T) {
	t.Run("normalize_path", func(t *testing.T) {
		for _, v := range tablenormalizePathInfo {
			n := newNode("get", v.in, func(w http.ResponseWriter, r *http.Request) {})

			if n.path != v.out {
				t.Errorf("Expected %s, but got: %s", v.out, n.path)
			}
		}
	})

	t.Run("normalize_method", func(t *testing.T) {
		for _, v := range tableNewNodeNotEmptyExistingMethodNormalize {
			n := newNode(v.in, "/", func(w http.ResponseWriter, r *http.Request) {})

			if n.method != v.out {
				t.Errorf("Expected %s to equal %s", v.out, n.method)
			}
		}
	})
}

func TestIsItForMe(t *testing.T) {
	t.Run("nil_request", func(t *testing.T) {
		for _, v := range tableIsItForMe {
			n := newNode(v.method, v.path, v.fnhandler)

			r := n.isItForMe(v.req)

			if v.want != r {
				t.Errorf("Request should be resolved %t, but got %t", v.want, r)
			}
		}
	})
}
