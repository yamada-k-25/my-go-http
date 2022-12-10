package main

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	crt, err := ioutil.ReadFile("../server/localhost.crt")
	if err != nil {
		log.Fatal(err)
	}
	block, _ := pem.Decode(crt)
	cert, err := x509.ParseCertificate(block.Bytes)

	// certPool := x509.NewCertPool()
	certPool, err := x509.SystemCertPool()
	if err != nil {
		log.Fatal(err)
	}
	certPool.AddCert(cert)

	client := http.Client{
		Transport: &http.Transport{},
	}

	// resp, err := client.Get("https://localhost:3000")
	resp, err := client.Post("https://localhost:3000", "application/json", bytes.NewReader([]byte("{\"message\":\"hello\"}")))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	file, err := os.Create("image.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(file, resp.Body)
	log.Printf("Protocol Version: %s\n", resp.Proto)
}
