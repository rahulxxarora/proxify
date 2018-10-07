package controllers

import (
	"fmt"
	"io"
	"net/http"

	"proxify/httpservice"
)

func HandleHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling HTTP request to %s\n", r.RequestURI)

	httpservice.MakeProxyHTTPReqest(w, r)
}

func HandleHTTPS(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling HTTPS request to %s\n", r.RequestURI)

	dest_conn := httpservice.SetConnectionToDestination(r)
	w.WriteHeader(http.StatusOK)

	hijacker, ok := w.(http.Hijacker)
	if !ok {
		panic("Hijacking not supported")
	}

	client_conn, _, err := hijacker.Hijack()
	if err != nil {
		panic(err)
	}

	go transfer(dest_conn, client_conn)
	go transfer(client_conn, dest_conn)
}

func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()

	io.Copy(destination, source)
}
