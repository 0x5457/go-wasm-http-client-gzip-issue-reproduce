package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"log"
	"net/http"
)

func httpNoGzip(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello, world")
}
func httpGzip(w http.ResponseWriter, req *http.Request) {
	var b bytes.Buffer
	gw := gzip.NewWriter(&b)
	gw.Write([]byte("hello, world"))
	gw.Close()

	w.Header().Set("Content-Encoding", "gzip")
	w.Write(b.Bytes())
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/gzip", httpGzip)
	http.HandleFunc("/nogzip", httpNoGzip)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
