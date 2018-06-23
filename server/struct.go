package klinserver

import (
	"sync"
)

type payload struct {
	C string `json:"content"`
	D bool   `json:"disabled"`
}

type payloadv2 struct {
	C string `json:"content"`
	D bool   `json:"disabled"`
}
type Dowork interface {
	Doit() string
}

func (m *payload) Doit() string {
	return m.C
}
func (m *payloadv2) Doit() string {
	return m.C + "noob"
}

type ServerConfig struct {
	Apikey   string
	Concur   int
	BindAddr string
	BindPort string
	Cert     string
	Key      string
	Trust    string
	Https    bool
}

type conn struct {
	regex   string
	apikey  string
	pkidir  string
	concur  int
	jobdir  string
	mu      sync.Mutex
	monorun chan struct{}
}
