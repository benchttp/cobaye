package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	port string

	requestCount int64
}

type urlpath string

const (
	urlpathDebug    urlpath = "/debug"
	urlpathIdentity urlpath = "/identity"
	urlpathMocks    urlpath = "/report"
)

func New(port string) *Server {
	return &Server{port: port}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch urlpath(r.URL.Path) {
	case urlpathDebug:
		s.handleDebug(w, r)
	case urlpathIdentity:
		s.handleIdentity(w, r)
	case urlpathMocks:
		s.handleMocks(w, r)
	default:
		s.handleFibonacci(w, r)
	}
}

func (s *Server) ListenAndServe() error {
	addr := "localhost:" + s.port
	fmt.Printf("http://%s\n", addr)

	return http.ListenAndServe(addr, s)
}
