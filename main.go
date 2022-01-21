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
	handler := http.HandlerFunc(handle)

	fmt.Printf("http://localhost:%s\n", *port)
	log.Fatal(http.ListenAndServe(addr, handler))
}
