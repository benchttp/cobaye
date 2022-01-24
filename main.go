package main

import (
	"flag"
	"log"

	"github.com/benchttp/cobaye/internal/server"
)

const defaultPort = "9999"

var port = flag.String("port", defaultPort, "listening port")

func main() {
	flag.Parse()
	s := server.New(*port)
	log.Fatal(s.ListenAndServe())
}
