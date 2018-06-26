package klinserver

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Server(c *ServerConfig) {
	certBytes, err := ioutil.ReadFile(c.Trust)
	if err != nil {
		log.Fatalln("Unable to read crt", err)
	}

	clientCertPool := x509.NewCertPool()
	if ok := clientCertPool.AppendCertsFromPEM(certBytes); !ok {
		log.Fatalln("Unable to add certificate to certificate pool")
	}

	tlsconfig := &tls.Config{
		PreferServerCipherSuites: true,
		// Only use curves which have assembly implementations
		ClientCAs: clientCertPool,
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519,
		},
		MinVersion: tls.VersionTLS12,
	}
	if c.Verify {
		tlsconfig.ClientAuth = tls.RequireAndVerifyClientCert
	}
	s := &http.Server{
		Addr:         c.BindAddr + ":" + c.BindPort,
		TLSConfig:    tlsconfig,
		Handler:      c.ServeMux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	fmt.Println("listening to " + c.BindAddr + " " + c.BindPort)
	if c.Https {
		err = s.ListenAndServeTLS(c.Cert, c.Key)
	} else {
		err = s.ListenAndServe()
	}
	if err != nil {
		log.Fatal("can't listen and serve check port and binding addr", err)
	}
}
