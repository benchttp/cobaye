package main

import (
	"flag"
	"log"

	"github.com/benchttp/cobaye/internal/server"
)

const defaultPort = "9999"

var (
	port        = flag.String("port", defaultPort, "listening port")
	ignoreStdin = flag.Bool("ignoreStdin", false, "do not listen stdin while serving")
)

func main() {
	flag.Parse()
	s := server.New(*port, *ignoreStdin)
	log.Fatal(s.ListenAndServe())
}
