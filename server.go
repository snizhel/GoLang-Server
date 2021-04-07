package main

import (
	"Server.go/router"
)

func main() {
	// e := router.Index()
	e := router.Routes()
	e.Start(":3000")
}
