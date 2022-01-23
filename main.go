package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const defaultPort = "9999"

var port = flag.String("port", defaultPort, "listening port")

func main() {
	flag.Parse()

	addr := ":" + *port
	handler := http.HandlerFunc(handleMain)

	fmt.Printf("http://localhost:%s\n", *port)
	go func() {
		log.Fatal(http.ListenAndServe(addr, handler))
	}()

	log.Fatal(listenInput())
}

func listenInput() error {
	reader := bufio.NewReader(os.Stdin)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("yo")
			return err
		}

		if string(line) == "debug" {
			fmt.Printf("Total requests: %d\n", totalRequests)
		}
	}
}
