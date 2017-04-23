package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	cert, err := tls.LoadX509KeyPair("client.pem", "client.key")
	if err != nil {
		fmt.Printf("#Error LoadX509KeyPair: %v\n", err.Error())
		return
	}

	certBytes, err := ioutil.ReadFile("client.pem")
	if err != nil {
		panic("Unable to read cert.pem")
	}
	clientCertPool := x509.NewCertPool()
	ok := clientCertPool.AppendCertsFromPEM(certBytes)
	if !ok {
		panic("Failed to parse root certificate")
	}

	conf := &tls.Config{
		RootCAs:            clientCertPool,
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", "127.0.0.1:443", conf)
	if err != nil {
		fmt.Printf("#Error Dial: %v\n", err.Error())
		return
	}
	defer conn.Close()

	for i := 0; i < 10; i++ {
		n, err := conn.Write([]byte("Hello\n"))
		if err != nil {
			fmt.Printf("#Error Write: %v, n=%v\n", err.Error(), n)
			return
		}

		buf := make([]byte, 100)
		n, err = conn.Read(buf)
		if err != nil {
			fmt.Printf("#Error Read: %v, n=%v\n", err.Error(), n)
			return
		}
		println(string(buf[:n]))
		time.Sleep(1*time.Second)
	}
}
