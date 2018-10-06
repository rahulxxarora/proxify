package app

import (
	"crypto/tls"
	"fmt"
	"net/http"

	. "proxify/controllers"
	_ "proxify/httpservice"
	. "proxify/middlewares"
)

type App struct {
	addr string
	port string
}

func (a *App) Initialize(addr string, port string) {
	fmt.Println("Initializing app...")

	a.addr = addr
	a.port = port

	fmt.Println("App initialized...")
}

func (a *App) Run() {
	fmt.Printf("HTTP server listening on %s:%s...\n", a.addr, a.port)

	server := &http.Server{
		Addr: a.addr + ":" + a.port,
		Handler: http.HandlerFunc(ErrorHandler(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodConnect {
				HandleHTTPS(w, r)
			} else {
				HandleHTTP(w, r)
			}
		})),
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}

	if error := server.ListenAndServe(); error != nil {
		panic(error)
	}
}
