package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	pool := x509.NewCertPool()
	caCertPath := filepath.Dir(os.Args[0]) + "/" + "ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	cliCrt, err := tls.LoadX509KeyPair(filepath.Dir(os.Args[0])+"/"+"client.crt", filepath.Dir(os.Args[0])+"/"+"client.key")
	if err != nil {
		fmt.Println("Loadx509keypair err:", err)
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt},
			ServerName:   "openpitrix.io",
		},
	}
	client := &http.Client{Transport: tr}
	//这里的ip地址需要在生成自签名证书的时候指定,否则ssl验证不通过。
	resp, err := client.Get("https://127.0.0.1:8081")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
