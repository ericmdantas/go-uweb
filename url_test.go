package uweb

import "testing"

var tablenormalizePathInfo = []struct {
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

func TestNormalizePath(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		for _, v := range tablenormalizePathInfo {
			r := normalizePath(v.in)

			if v.out != r {
				t.Errorf("Expected %s but got %s", v.out, r)
			}
		}
	})
}

func BenchmarNormalizePath(b *testing.B) {
	b.Run("empty_string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			normalizePath("")
		}
	})

	b.Run("only_slash", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			normalizePath("/")
		}
	})

	b.Run("slash_ending", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			normalizePath("/abc/")
		}
	})

	b.Run("correct_slash", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			normalizePath("/api/todos/:id")
		}
	})

	b.Run("long", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			normalizePath("/api/todos/:id/:name/:something/:this/:that/:action/:yo")
		}
	})
}
