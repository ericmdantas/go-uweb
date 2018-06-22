package uweb

import (
	"net/http"
	"testing"
)

var tableTestAddNodesMethods = []struct {
	method    string
	path      string
	mapkey    string
	fnhandler UWebHandlerFunc
	want      *Node
}{
	{
		method:    "GET",
		path:      "/hello_world",
		mapkey:    "GET /hello_world",
		fnhandler: func(w http.ResponseWriter, r *http.Request) {},
		want: &Node{
			method: "GET",
			path:   "/hello_world",
			handle: func(w http.ResponseWriter, r *http.Request) {},
		},
	},
	{
		method:    "POST",
		path:      "/hello_world",
		mapkey:    "POST /hello_world",
		fnhandler: func(w http.ResponseWriter, r *http.Request) {},
		want: &Node{
			method: "POST",
			path:   "/hello_world",
			handle: func(w http.ResponseWriter, r *http.Request) {},
		},
	},
	{
		method:    "PUT",
		path:      "/hello_world",
		mapkey:    "PUT /hello_world",
		fnhandler: func(w http.ResponseWriter, r *http.Request) {},
		want: &Node{
			method: "PUT",
			path:   "/hello_world",
			handle: func(w http.ResponseWriter, r *http.Request) {},
		},
	},
	{
		method:    "HEAD",
		path:      "/hello_world",
		mapkey:    "HEAD /hello_world",
		fnhandler: func(w http.ResponseWriter, r *http.Request) {},
		want: &Node{
			method: "HEAD",
			path:   "/hello_world",
			handle: func(w http.ResponseWriter, r *http.Request) {},
		},
	},
	{
		method:    "TRACE",
		path:      "/hello_world",
		mapkey:    "TRACE /hello_world",
		fnhandler: func(w http.ResponseWriter, r *http.Request) {},
		want: &Node{
			method: "TRACE",
			path:   "/hello_world",
			handle: func(w http.ResponseWriter, r *http.Request) {},
		},
	},
	{
		method:    "CONNECT",
		path:      "/hello_world",
		mapkey:    "CONNECT /hello_world",
		fnhandler: func(w http.ResponseWriter, r *http.Request) {},
		want: &Node{
			method: "CONNECT",
			path:   "/hello_world",
			handle: func(w http.ResponseWriter, r *http.Request) {},
		},
	},
}

func TestMethods(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		for _, v := range tableTestAddNodesMethods {
			u := New()

			switch v.method {
			case GET:
				u.GET(v.path, v.fnhandler)
				break
			case POST:
				u.POST(v.path, v.fnhandler)
				break
			case PUT:
				u.PUT(v.path, v.fnhandler)
				break
			case HEAD:
				u.HEAD(v.path, v.fnhandler)
				break
			case TRACE:
				u.TRACE(v.path, v.fnhandler)
				break
			case CONNECT:
				u.CONNECT(v.path, v.fnhandler)
				break
			case DELETE:
				u.DELETE(v.path, v.fnhandler)
				break
			}

			if u.tree[v.mapkey].method != v.want.method {
				t.Errorf("Different method. Want %s, but got %s", v.want.method, u.tree[v.mapkey].method)
			}

			if u.tree[v.mapkey].path != v.want.path {
				t.Errorf("Different path. Want %s, but got %s", v.want.path, u.tree[v.mapkey].path)
			}
		}
	})
}

func BenchmarkMethods(b *testing.B) {
	b.Run("GET", func(b *testing.B) {
		u := New()
		msg := []byte("!")

		for i := 0; i < b.N; i++ {
			u.GET("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write(msg)
			})
		}
	})

	b.Run("POST", func(b *testing.B) {
		u := New()
		msg := []byte("!")

		for i := 0; i < b.N; i++ {
			u.POST("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write(msg)
			})
		}
	})

	b.Run("PUT", func(b *testing.B) {
		u := New()
		msg := []byte("!")

		for i := 0; i < b.N; i++ {
			u.PUT("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write(msg)
			})
		}
	})

	b.Run("PATCH", func(b *testing.B) {
		u := New()
		msg := []byte("!")

		for i := 0; i < b.N; i++ {
			u.PATCH("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write(msg)
			})
		}
	})

	b.Run("HEAD", func(b *testing.B) {
		u := New()
		msg := []byte("!")

		for i := 0; i < b.N; i++ {
			u.HEAD("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write(msg)
			})
		}
	})

	b.Run("DELETE", func(b *testing.B) {
		u := New()
		msg := []byte("!")

		for i := 0; i < b.N; i++ {
			u.DELETE("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write(msg)
			})
		}
	})

	b.Run("OPTIONS", func(b *testing.B) {
		u := New()
		msg := []byte("!")

		for i := 0; i < b.N; i++ {
			u.OPTIONS("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write(msg)
			})
		}
	})

	b.Run("CONNECT", func(b *testing.B) {
		u := New()
		msg := []byte("!")

		for i := 0; i < b.N; i++ {
			u.CONNECT("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write(msg)
			})
		}
	})

	b.Run("TRACE", func(b *testing.B) {
		u := New()
		msg := []byte("!")

		for i := 0; i < b.N; i++ {
			u.TRACE("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write(msg)
			})
		}
	})
}
