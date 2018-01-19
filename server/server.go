package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type myhandler struct {
}

func (h *myhandler) ServeHTTP(w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprintf(w,
		"Hi, This is an example of http service in golang!\n")
}

func main() {
	pool := x509.NewCertPool()
	caCertPath := filepath.Dir(os.Args[0]) + "/" + "ca.crt"
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)
	//初始化一个server 实例。
	s := &http.Server{
		//设置宿主机的ip地址，并且端口号为8081
		Addr:    ":8081",
		Handler: &myhandler{},
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	err = s.ListenAndServeTLS(filepath.Dir(os.Args[0])+"/"+"server.crt", filepath.Dir(os.Args[0])+"/"+"server.key")

	if err != nil {
		fmt.Println("ListenAndServeTLS err:", err)
	}
}
