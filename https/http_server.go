package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

type myHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "\nHi, This is an example of http service in golang!\n")
}

func main() {
	pool := x509.NewCertPool()
	caCrt, err := ioutil.ReadFile("ca.crt")
	checkErrors(err)
	pool.AppendCertsFromPEM(caCrt)

	server := &http.Server{
		Addr:    ":8082",
		Handler: &myHandler{},
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	logrus.Error(server.ListenAndServeTLS("server.crt", "server.key"))
}

func checkErrors(err error) {
	if err != nil {
		fmt.Printf("#Error: %v\n", err.Error())
		return
	}
}
