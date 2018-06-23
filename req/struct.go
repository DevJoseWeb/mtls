package klinreq

import ()

type ReqInfo struct {
	Cert   string // The cert for mtls
	Key    string // The key for mtls
	Dest   string // The destination address. It has to be hostname
	Dport  string // The destination address port
	Trust  string // The trusted CA cert
	Method string // The req method, POST/PATCH etc...
	Route  string // The route, by default its "/" it can be "/api"
	File   string // If you are sending file specify the file you are sending.
}
