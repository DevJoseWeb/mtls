package klinserver

import (
	"fmt"
	"net/http"
	"testing"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
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
		BindPort: "2018",
		Cert:     "program/test3.klin-pro.com.crt",
		Key:      "program/test3.klin-pro.com.key",
		Trust:    "program/mtls.crt",
		Https:    true,
		Verify:   true,
		ServeMux: con,
	}
	Server(s)
}
