package main

import (
	"flag"
	"log"

	"github.com/benchttp/cobaye/server"
)

const defaultPort = "9999"

var port = flag.String("port", defaultPort, "listening port")

func main() {
	flag.Parse()
	s := server.New(*port)

	go s.ListenStdin()

	log.Fatal(s.ListenAndServe())
}
