package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080/gzip")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Content-Encoding:", resp.Header.Get("Content-Encoding"))
	fmt.Println("Content-Length:", resp.Header.Get("Content-Length"))
	fmt.Println("ContentLength:", resp.ContentLength)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("body:", body)
}
