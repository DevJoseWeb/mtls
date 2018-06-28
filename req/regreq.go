package klinreq

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// Send a json payload. payload should be a struct where you define your json
func SendPayload(i *ReqInfo, payload interface{}) (*http.Response, error) {
	var resp *http.Response
	cert, err := tls.LoadX509KeyPair(i.Cert, i.Key)
	if err != nil {
		return resp, err
	}

	// Load our CA certificate
	clientCACert, err := ioutil.ReadFile(i.Trust)
	if err != nil {
		return resp, err
	}

	clientCertPool := x509.NewCertPool()
	clientCertPool.AppendCertsFromPEM(clientCACert)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		Certificates:       []tls.Certificate{cert},
		RootCAs:            clientCertPool,
	}
	tr := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{
		Timeout:   500 * time.Millisecond,
		Transport: tr,
	}

	encodepayload, _ := json.Marshal(payload)
	ebody := bytes.NewReader(encodepayload)
	var addr string
	if i.Http {
		addr = "http://" + i.Dest + ":" + i.Dport + "/" + i.Route
	} else {
		addr = "https://" + i.Dest + ":" + i.Dport + "/" + i.Route
	}
	req, err := http.NewRequest(i.Method, addr, ebody)
	if err != nil {
		return resp, err
	}
	resp, err = client.Do(req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
