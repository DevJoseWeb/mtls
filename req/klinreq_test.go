package klinreq

import (
	"fmt"
	"io/ioutil"
	"testing"
)

type testPayload struct {
	C string `json:"content"`
	D bool   `json:"disabled"`
}

func TestReq(t *testing.T) {
	fmt.Println("testing req")
	payload := &testPayload{
		C: "wtf",
		D: true,
	}
	i := &ReqInfo{
		Cert:    "program/test2.klin-pro.com.crt",
		Key:     "program/test2.klin-pro.com.key",
		Dest:    "test3.klin-pro.com",
		Dport:   "2018",
		Trust:   "program/rootca.crt",
		Method:  "POST",
		Route:   "foo",
		Payload: payload,
	}
	resp, err := SendPayload(i)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body), string(resp.Status))
}

func TestSendFile(t *testing.T) {
	fmt.Println("testing filesend")
	i := &ReqInfo{
		Cert:   "program/test2.klin-pro.com.crt",
		Key:    "program/test2.klin-pro.com.key",
		Dest:   "test3.klin-pro.com",
		Dport:  "2018",
		Trust:  "program/rootca.crt",
		Method: "POST",
		File:   "program/testfile",
		Route:  "foo",
		ExtraParams: map[string]string{
			"filename": "klinFile",
		},
	}
	resp := SendFile(i)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body), string(resp.Status))
}
