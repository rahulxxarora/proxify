package app

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strconv"

	. "github.com/rahulxxarora/proxify/controllers"
	_ "github.com/rahulxxarora/proxify/httpservice"
	. "github.com/rahulxxarora/proxify/middlewares"
)

type App struct {
	addr string
	port int
}

func (a *App) Initialize(addr string, port int) {
	fmt.Println("Initializing app...")

	a.addr = addr
	a.port = port

	fmt.Println("App initialized...")
}

func (a *App) Run() {
	fmt.Printf("HTTP server listening on %s:%d...\n", a.addr, a.port)

	server := &http.Server{
		Addr: a.addr + ":" + strconv.Itoa(a.port),
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
