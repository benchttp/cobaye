package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/benchttp/cobaye/internal/server"
)

const defaultPort = "9999"

var port = flag.String("port", defaultPort, "listening port")

func main() {
	flag.Parse()

	addr := ":" + *port
	s := server.Server{}
	handler := http.HandlerFunc(s.HandleMain)

	fmt.Printf("http://localhost%s\n", addr)
	go func() {
		log.Fatal(http.ListenAndServe(addr, handler))
	}()

	log.Fatal(s.ListenInput())
}
