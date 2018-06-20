package uweb

import (
	"net/http"
	"testing"
)

var tableTestAddNodes = []struct {
	method    string
	path      string
	fnhandler UWebHandlerFunc
	want      []*Node
}{
	{
		method:    "GET",
		path:      "/hello_world",
		fnhandler: func(rw http.ResponseWriter, rq *http.Request) {},
		want: []*Node{
			{
				method: "GET",
				path:   "/hello_world",
				handle: func(rw http.ResponseWriter, rq *http.Request) {},
			},
		},
	},
}

func TestAddNode(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		for _, v := range tableTestAddNodes {
			u := New()

			u.addNode(v.method, v.path, v.fnhandler)

			if u.tree[0].method != v.want[0].method {
				t.Errorf("Different method. Want %s, but got %s", v.want[0].method, u.tree[0].method)
			}

			if u.tree[0].path != v.want[0].path {
				t.Errorf("Different path. Want %s, but got %s", v.want[0].path, u.tree[0].path)
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
