package uweb

import (
	"net/http"
	"testing"
)

func BenchmarkGETSimple(b *testing.B) {
	u := New()
	msg := []byte("!")

	for i := 0; i < b.N; i++ {
		u.GET("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(msg)
		})
	}
}

func BenchmarkPOSTSimple(b *testing.B) {
	u := New()
	msg := []byte("!")

	for i := 0; i < b.N; i++ {
		u.POST("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(msg)
		})
	}
}

func BenchmarkPUTSimple(b *testing.B) {
	u := New()
	msg := []byte("!")

	for i := 0; i < b.N; i++ {
		u.PUT("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(msg)
		})
	}
}

func BenchmarkPATCHSimple(b *testing.B) {
	u := New()
	msg := []byte("!")

	for i := 0; i < b.N; i++ {
		u.PATCH("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(msg)
		})
	}
}

func BenchmarkHEADSimple(b *testing.B) {
	u := New()
	msg := []byte("!")

	for i := 0; i < b.N; i++ {
		u.HEAD("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(msg)
		})
	}
}

func BenchmarkDELETESimple(b *testing.B) {
	u := New()
	msg := []byte("!")

	for i := 0; i < b.N; i++ {
		u.DELETE("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(msg)
		})
	}
}

func BenchmarkOPTIONSSimple(b *testing.B) {
	u := New()
	msg := []byte("!")

	for i := 0; i < b.N; i++ {
		u.OPTIONS("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(msg)
		})
	}
}

func BenchmarkCONNECTSimple(b *testing.B) {
	u := New()
	msg := []byte("!")

	for i := 0; i < b.N; i++ {
		u.CONNECT("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(msg)
		})
	}
}

func BenchmarkTRACESimple(b *testing.B) {
	u := New()
	msg := []byte("!")

	for i := 0; i < b.N; i++ {
		u.TRACE("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(msg)
		})
	}
}
