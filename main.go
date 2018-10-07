package main

import (
	"flag"
	. "proxify/app"
)

func main() {
	addr := flag.String("proxy-host", "127.0.0.1", "host address")
	port := flag.Int("proxy-port", 9090, "host port")
	flag.Parse()

	a := App{}
	a.Initialize(*addr, *port)
	a.Run()
}
