package server

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync/atomic"
)

type Server struct {
	port string

	requestCount int64
}

func New(port string) *Server {
	return &Server{port: port}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.handle(w, r)
}

func (s *Server) ListenAndServe() error {
	addr := "localhost:" + s.port
	fmt.Printf("http://%s\n", addr)
	go s.listenStdin()
	return http.ListenAndServe(addr, s) //nolint:gosec // local use only
}

func (s *Server) listenStdin() {
	reader := bufio.NewReader(os.Stdin)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println(err) // non critical
		}

		if string(line) == "debug" {
			fmt.Printf("Total requests: %d\n", s.requestCount)
		}
	}
}

func (s *Server) incrementRequestCount() {
	atomic.AddInt64(&s.requestCount, 1)
}
