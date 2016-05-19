package uweb

import "testing"

var tableNormalizePathInfo = []struct {
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
	for _, v := range tableNormalizePathInfo {
		r := NormalizePath(v.in)

		if v.out != r {
			t.Errorf("Expected %s but got %s", v.out, r)
		}
	}
}

func BenchmarkNormalizePathEmptyString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NormalizePath("")
	}
}

func BenchmarkNormalizePathOnlySlash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NormalizePath("/")
	}
}

func BenchmarkNormalizePathSlashEnding(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NormalizePath("/abc/")
	}
}

func BenchmarkNormalizePathCorrectSlash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NormalizePath("/api/todos/:id")
	}
}

func BenchmarkNormalizePathLong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NormalizePath("/api/todos/:id/:name/:something/:this/:that/:action/:yo")
	}
}
