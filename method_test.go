package uweb

import (
	"testing"
)

var tableNotEmptyExistingMethodNormalize = []struct {
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

var tableExistingMethodWithSillyErrorsNormalize = []struct {
	in, out string
}{
	{
		in:  "get    ",
		out: "GET",
	},
	{
		in:  "     head     ",
		out: "HEAD",
	},
	{
		in:  "   post ",
		out: "POST",
	},
	{
		in:  "pUt",
		out: "PUT",
	},
	{
		in:  "opTIons ",
		out: "OPTIONS",
	},
	{
		in:  "deLete",
		out: "DELETE",
	},
	{
		in:  "connecT",
		out: "CONNECT",
	},
	{
		in:  "tracE",
		out: "TRACE",
	},
	{
		in:  "patcH",
		out: "PATCH",
	},
}

var tableNotExistingMethods = []string{
	"",
	"git",
	"headi",
	"hea",
	"heas",
	"putz",
}

func BenchmarkNormalizeMethod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		normalizeMethod("get    ")
	}
}

func TestNormalizeExistingMethods(t *testing.T) {
	for _, v := range tableNotEmptyExistingMethodNormalize {
		r := normalizeMethod(v.in)

		if r != v.out {
			t.Errorf("Expected %s to equal %s", v.out, r)
		}
	}
}

func TestNormalizeExistingMethodsWithSillyErrors(t *testing.T) {
	for _, v := range tableExistingMethodWithSillyErrorsNormalize {
		r := normalizeMethod(v.in)

		if r != v.out {
			t.Errorf("Expected %s, but got: %s", v.out, r)
		}
	}
}

func TestNormalizeNotExistingMethods(t *testing.T) {
	for _, v := range tableNotExistingMethods {
		var err interface{}

		defer func() {
			if err == nil {
				t.Errorf("Should've paniced")
			}
		}()

		defer func() {
			err = recover()
		}()

		normalizeMethod(v)
	}
}
