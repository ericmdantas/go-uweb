package uweb

import (
	"net/url"
	"testing"
)

func BenchmarkParseSimple(b *testing.B) {
	p := Params{}

	for i := 0; i < b.N; i++ {
		p.Parse("/")
	}
}

func TestParseSimpleSingleSlash(t *testing.T) {
	p := Params{}
	u := "/"

	p.Parse(u)

	if p.raw != u {
		t.Errorf("expected raw to equal %s, but got %s", u, p.raw)
	}
}

func TestParseWithParams(t *testing.T) {
	p := Params{}
	u := "/:id"
	parsed, _ := url.Parse(u)

	p.Parse(u)

	if p.raw != u {
		t.Errorf("expected raw to equal %s, but got %s", u, p.raw)
	}

	if p.parsed.String() != parsed.String() {
		t.Errorf("expected parsed to equal %s, but got %s", u, p.parsed)
	}
}

func TestParseWithQueryString(t *testing.T) {
	p := Params{}
	u := "/?abc=123"
	parsed, _ := url.Parse(u)

	p.Parse(u)

	if p.raw != u {
		t.Errorf("expected raw to equal %s, but got %s", u, p.raw)
	}

	if p.parsed.String() != parsed.String() {
		t.Errorf("expected parsed to equal %s, but got %s", u, p.parsed)
	}
}

func TestParseWithParamsAndQueryString(t *testing.T) {
	p := Params{}
	u := "/:id?abc=123"
	parsed, _ := url.Parse(u)

	p.Parse(u)

	if p.raw != u {
		t.Errorf("expected raw to equal %s, but got %s", u, p.raw)
	}

	if p.parsed.String() != parsed.String() {
		t.Errorf("expected parsed to equal %s, but got %s", u, p.parsed)
	}
}
