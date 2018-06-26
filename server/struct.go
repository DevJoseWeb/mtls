package klinserver

import (
	"net/http"
)

type ServerConfig struct {
	BindAddr string
	BindPort string
	Cert     string
	Key      string
	Trust    string
	Https    bool
	Verify   bool
	ServeMux *http.ServeMux
}
