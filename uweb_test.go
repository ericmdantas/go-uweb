package uweb

import (
	"net/http"
	"testing"
)

func BenchmarkGetSimple(b *testing.B) {
	u := New()
	msg := []byte("!")

	for i := 0; i < b.N; i++ {
		u.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(msg)
		})
	}
}

func BenchmarkPostSimple(b *testing.B) {
	u := New()
	msg := []byte("!")

	for i := 0; i < b.N; i++ {
		u.Post("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(msg)
		})
	}
}

func BenchmarkPutSimple(b *testing.B) {
	u := New()
	msg := []byte("!")

	for i := 0; i < b.N; i++ {
		u.Put("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(msg)
		})
	}
}

func BenchmarkPatchSimple(b *testing.B) {
	u := New()
	msg := []byte("!")

	for i := 0; i < b.N; i++ {
		u.Patch("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(msg)
		})
	}
}

func BenchmarkDeleteSimple(b *testing.B) {
	u := New()
	msg := []byte("!")

	for i := 0; i < b.N; i++ {
		u.Delete("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(msg)
		})
	}
}

func BenchmarkOptionsSimple(b *testing.B) {
	u := New()
	msg := []byte("!")

	for i := 0; i < b.N; i++ {
		u.Options("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(msg)
		})
	}
}

func BenchmarkConnectSimple(b *testing.B) {
	u := New()
	msg := []byte("!")

	for i := 0; i < b.N; i++ {
		u.Connect("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(msg)
		})
	}
}

func BenchmarkTraceSimple(b *testing.B) {
	u := New()
	msg := []byte("!")

	for i := 0; i < b.N; i++ {
		u.Connect("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(msg)
		})
	}
}
