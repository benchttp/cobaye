package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const defaultPort = "9999"

var port = flag.String("port", defaultPort, "listening port")

func main() {
	flag.Parse()

	addr := ":" + *port
	server := server{}
	handler := http.HandlerFunc(server.handleMain)

	fmt.Printf("http://localhost%s\n", addr)
	go func() {
		log.Fatal(http.ListenAndServe(addr, handler))
	}()

	log.Fatal(server.listenInput())
}
