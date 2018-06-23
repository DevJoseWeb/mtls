package klinserver

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	fmt.Println("testing server")
	c := new(conn)
	sema := make(chan struct{}, 1)
	c.monorun = sema
	con := http.NewServeMux()
	con.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var body payload
		c.notwork(w, r, &body)
	})
	con.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		var body payloadv2
		c.handleWebHook(w, r, &body)
	})
	s := &ServerConfig{
		Apikey:   "wtf",
		Concur:   5,
		BindPort: "2018",
		Cert:     "test3.crt",
		Key:      "test3.key",
		Trust:    "devca.crt",
		Https:    true,
		ServeMux: con,
	}
	Server(s)
}
