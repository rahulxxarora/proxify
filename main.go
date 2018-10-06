package main

import (
	. "proxify/app"
)

func main() {
	a := App{}
	a.Initialize("127.0.0.1", "9090")
	a.Run()
}
