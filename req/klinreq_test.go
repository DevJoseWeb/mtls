package klinreq

import (
	"fmt"
	"testing"
)

type testPayload struct {
	C string `json:"content"`
	D bool   `json:"disabled"`
}

func TestReq(t *testing.T) {
	fmt.Println("testing req")
	i := &ReqInfo{
		Cert:   "program/test2.klin-pro.com.crt",
		Key:    "program/test2.klin-pro.com.key",
		Dest:   "test3.klin-pro.com",
		Dport:  "2018",
		Trust:  "program/master.pem",
		Method: "POST",
		Route:  "nigga/shji",
	}
	payload := &testPayload{
		C: "wtf",
		D: true,
	}
	SendPayload(i, payload)
}

func TestSendfile(t *testing.T) {
	fmt.Println("testing filesend")
	i := &ReqInfo{
		Cert:   "program/test2.klin-pro.com.crt",
		Key:    "program/test2.klin-pro.com.key",
		Dest:   "test3.klin-pro.com",
		Dport:  "2018",
		Trust:  "program/master.pem",
		Method: "POST",
		File:   "program/testfile",
		Route:  "foo",
	}
	Sendfile(i)
}
