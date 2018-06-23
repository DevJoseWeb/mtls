package klinserver

import (
	"sync"
)

type ServerConfig struct {
	Apikey   string
	Concur   int
	BindAddr string
	BindPort string
	Cert     string
	Key      string
	Trust    string
	Https    bool
	Payload  interface{}
}

type conn struct {
	regex   string
	apikey  string
	pkidir  string
	concur  int
	jobdir  string
	mu      sync.Mutex
	monorun chan struct{}
	payload interface{}
}
