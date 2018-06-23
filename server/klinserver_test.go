package klinserver

import (
	"fmt"
	"testing"
)

func TestServer(t *testing.T) {
	fmt.Println("testing server")
	s := &ServerConfig{
		Apikey:   "wtf",
		Concur:   5,
		BindPort: "2018",
		Cert:     "test3.crt",
		Key:      "test3.key",
		Trust:    "devca.crt",
		Https:    true,
	}
	Server(s)
}
