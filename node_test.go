package uweb

import (
	"net/url"
	"testing"
)

func BenchmarkParseSimple(b *testing.B) {
	n := Node{}

	for i := 0; i < b.N; i++ {
		n.Parse()
	}
}

func TestParseSimpleSingleSlash(t *testing.T) {
	u := "/"
	n := Node{path: "/"}

	n.Parse()

	if n.path != u {
		t.Errorf("expected path to equal %s, but got %s", u, n.path)
	}
}

func TestParseWithParam(t *testing.T) {
	u := "/:id"
	n := Node{path: "/:id"}
	parsed, _ := url.Parse(u)

	n.Parse()

	if n.path != u {
		t.Errorf("expected path to equal %s, but got %s", u, n.path)
	}

	if n.parsed.String() != parsed.String() {
		t.Errorf("expected parsed to equal %s, but got %s", u, n.parsed)
	}
}

func TestParseWithQueryString(t *testing.T) {
	u := "/?abc=123"
	n := Node{path: "/"}
	parsed, _ := url.Parse(u)

	n.Parse()

	if n.path != u {
		t.Errorf("expected path to equal %s, but got %s", u, n.path)
	}

	if n.parsed.String() != parsed.String() {
		t.Errorf("expected parsed to equal %s, but got %s", u, n.parsed)
	}
}

func TestParseWithParamAndQueryString(t *testing.T) {
	u := "/:id?abc=123"
	n := Node{path: "/:id"}
	parsed, _ := url.Parse(u)

	n.Parse()

	if n.path != u {
		t.Errorf("expected path to equal %s, but got %s", u, n.path)
	}

	if n.parsed.String() != parsed.String() {
		t.Errorf("expected parsed to equal %s, but got %s", u, n.parsed)
	}
}
