package uweb

import (
	"net/http"
	"testing"
)

func BenchmarkGET(b *testing.B) {
	b.Run("simple", func(b *testing.B) {
		u := New()
		msg := []byte("!")

		for i := 0; i < b.N; i++ {
			u.GET("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write(msg)
			})
		}
	})
}

func BenchmarkPOST(b *testing.B) {
	b.Run("simple", func(b *testing.B) {
		u := New()
		msg := []byte("!")

		for i := 0; i < b.N; i++ {
			u.POST("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write(msg)
			})
		}
	})
}

func BenchmarkPUT(b *testing.B) {
	b.Run("simple", func(b *testing.B) {
		u := New()
		msg := []byte("!")

		for i := 0; i < b.N; i++ {
			u.PUT("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write(msg)
			})
		}
	})
}

func BenchmarkPATCH(b *testing.B) {
	b.Run("simple", func(b *testing.B) {
		u := New()
		msg := []byte("!")

		for i := 0; i < b.N; i++ {
			u.PATCH("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write(msg)
			})
		}
	})
}

func BenchmarkHEAD(b *testing.B) {
	b.Run("simple", func(b *testing.B) {
		u := New()
		msg := []byte("!")

		for i := 0; i < b.N; i++ {
			u.HEAD("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write(msg)
			})
		}
	})
}

func BenchmarkDELETE(b *testing.B) {
	b.Run("simple", func(b *testing.B) {
		u := New()
		msg := []byte("!")

		for i := 0; i < b.N; i++ {
			u.DELETE("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write(msg)
			})
		}
	})
}

func BenchmarkOPTIONS(b *testing.B) {
	b.Run("simple", func(b *testing.B) {
		u := New()
		msg := []byte("!")

		for i := 0; i < b.N; i++ {
			u.OPTIONS("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write(msg)
			})
		}
	})
}

func BenchmarkCONNECT(b *testing.B) {
	b.Run("simple", func(b *testing.B) {
		u := New()
		msg := []byte("!")

		for i := 0; i < b.N; i++ {
			u.CONNECT("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write(msg)
			})
		}
	})
}

func BenchmarkTRACE(b *testing.B) {
	b.Run("simple", func(b *testing.B) {
		u := New()
		msg := []byte("!")

		for i := 0; i < b.N; i++ {
			u.TRACE("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write(msg)
			})
		}
	})
}
