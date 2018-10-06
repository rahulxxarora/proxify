package controllers

import (
	"fmt"
	"net/http"

	"proxify/httpservice"
)

func HandleHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling HTTP request to %s\n", r.RequestURI)

	httpservice.MakeProxyHTTPReqest(w, r)
}

func HandleHTTPS(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling HTTPS request to %s\n", r.RequestURI)
}
