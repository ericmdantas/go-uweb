package uweb

import (
	"net/url"
)

type Params struct {
	raw    string
	parsed *url.URL
	path   map[string]string
	query  map[string][]string
}

func (p *Params) Parse(u string) {
	p.raw = u
	p.parsed, _ = url.Parse(u)
}

func (p *Params) Query(k string) map[string]string {
	var q map[string]string

	return q
}

func (p *Params) PathParams(k string) map[string]string {
	var q map[string]string

	return q
}
