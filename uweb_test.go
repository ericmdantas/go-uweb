package uweb

import (
	"net/http"
	"testing"
)

func BenchmarkGet(b *testing.B) {
	u := New()
	msg := []byte("!")

	for i := 0; i < b.N; i++ {
		u.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(msg)
		})
	}
}

func BenchmarkPost(b *testing.B) {
	u := New()
	msg := []byte("!")

	for i := 0; i < b.N; i++ {
		u.Post("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(msg)
		})
	}
}

func BenchmarkPut(b *testing.B) {
	u := New()
	msg := []byte("!")

	for i := 0; i < b.N; i++ {
		u.Put("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(msg)
		})
	}
}

func BenchmarkPatch(b *testing.B) {
	u := New()
	msg := []byte("!")

	for i := 0; i < b.N; i++ {
		u.Patch("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(msg)
		})
	}
}

func BenchmarkDelete(b *testing.B) {
	u := New()
	msg := []byte("!")

	for i := 0; i < b.N; i++ {
		u.Delete("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(msg)
		})
	}
}

func BenchmarkOptions(b *testing.B) {
	u := New()
	msg := []byte("!")

	for i := 0; i < b.N; i++ {
		u.Options("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(msg)
		})
	}
}

func BenchmarkConnect(b *testing.B) {
	u := New()
	msg := []byte("!")

	for i := 0; i < b.N; i++ {
		u.Connect("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(msg)
		})
	}
}

func BenchmarkTrace(b *testing.B) {
	u := New()
	msg := []byte("!")

	for i := 0; i < b.N; i++ {
		u.Connect("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(msg)
		})
	}
}
