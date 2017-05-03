package uweb

import (
	"net/http"
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

func BenchmarkNewNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newNode("Get", "/", func(w http.ResponseWriter, r *http.Request) {})
	}
}

func TestNewNodenormalizePath(t *testing.T) {
	for _, v := range tablenormalizePathInfo {
		n := newNode("get", v.in, func(w http.ResponseWriter, r *http.Request) {})

		if n.path != v.out {
			t.Errorf("Expected %s, but got: %s", v.out, n.path)
		}
	}
}

func TestNewNodenormalizeMethod(t *testing.T) {
	for _, v := range tableNewNodeNotEmptyExistingMethodNormalize {
		n := newNode(v.in, "/", func(w http.ResponseWriter, r *http.Request) {})

		if n.method != v.out {
			t.Errorf("Expected %s to equal %s", v.out, n.method)
		}
	}
}
