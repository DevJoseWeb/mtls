package klinserver

import (
	"fmt"
	"testing"
)

type testPayload struct {
	C string `json:"content"`
	D bool   `json:"disabled"`
}

func TestServer(t *testing.T) {
	fmt.Println("testing server")
	var p testPayload
	s := &ServerConfig{
		Apikey:   "wtf",
		Concur:   5,
		BindPort: "2018",
		Cert:     "test3.crt",
		Key:      "test3.key",
		Trust:    "devca.crt",
		Https:    true,
		Payload:  p,
	}
	Server(s)
}
