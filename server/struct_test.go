package klinserver

import (
	"sync"
)

type conn struct {
	regex   string
	apikey  string
	pkidir  string
	concur  int
	jobdir  string
	mu      *sync.Mutex
	monorun chan struct{}
}

type payload struct {
	C string `json:"content"`
	D bool   `json:"disabled"`
}

type payloadv2 struct {
	C string `json:"content"`
	D bool   `json:"disabled"`
}
type dowork interface {
	doit() string
}

func (m *payload) doit() string {
	return m.C
}
func (m *payloadv2) doit() string {
	return m.C + "noob"
}
