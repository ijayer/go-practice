package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 1. 客户端不做任何验证
	resp, err := http.Get("https://182.254.245.83:8080/v1.0/versions/all")
	checkError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	// 2. 客户端验证Server证书
	//pool := x509.NewCertPool()
	//caCrt, err := ioutil.ReadFile("ca.crt")
	//checkError(err)
	//pool.AppendCertsFromPEM(caCrt)
	//
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{RootCAs: pool},
	//}
	//client := &http.Client{Transport: tr}
	//
	//resp, err := client.Get("https://localhost/v1.0/versions/all")
	//checkError(err)
	//defer resp.Body.Close()
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

	// 3. 对客户端的证书进行校验(双向证书校验）
	//pool := x509.NewCertPool()
	//caCertPath := "ca.crt"
	//
	//caCrt, err := ioutil.ReadFile(caCertPath)
	//checkError(err)
	//pool.AppendCertsFromPEM(caCrt)
	//
	//cliCrt, err := tls.LoadX509KeyPair("client.crt", "client.key")
	//checkError(err)
	//
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{
	//		RootCAs:      pool,
	//		Certificates: []tls.Certificate{cliCrt},
	//	},
	//}
	//client := &http.Client{Transport: tr}
	//resp, err := client.Get("https://182.254.245.83:8080/v1.0/versions/all")
	//checkError(err)
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("#Error: %v\n", err.Error())
		return
	}
}
