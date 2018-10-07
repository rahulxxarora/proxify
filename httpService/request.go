package httpservice

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

var (
	httpClient *http.Client
)

func init() {
	fmt.Println("Initializing HTTP client...")

	httpClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 5,
		},
		Timeout: 10 * time.Second,
	}

	fmt.Println("HTTP client initialized...")
}

func MakeProxyHTTPReqest(w http.ResponseWriter, r *http.Request) {
	// Initialize a new reques, same as the client's request
	request, err := http.NewRequest(r.Method, r.RequestURI, r.Body)
	if err != nil {
		panic(err)
	}

	// Copy all the headers set by client
	for haderName, headerValue := range r.Header {
		request.Header.Set(haderName, headerValue[0])
	}

	// Make HTTP request
	response, err := httpClient.Do(request)
	if err != nil {
		panic(err)
	}

	r.Body.Close()

	// Copy all the response headers
	for headerName, headerValue := range response.Header {
		w.Header().Set(headerName, headerValue[0])
	}

	w.WriteHeader(response.StatusCode)
	io.Copy(w, response.Body)
	response.Body.Close()
}

func SetConnectionToDestination(r *http.Request) net.Conn {
	dest_conn, err := net.DialTimeout("tcp", r.Host, 10*time.Second)

	if err != nil {
		panic(err)
	}

	return dest_conn
}
