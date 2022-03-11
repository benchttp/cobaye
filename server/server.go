package server

import (
	"fmt"
	"net/http"
	"path"
	"strings"
)

type Server struct {
	port string

	requestCount int64
}

func New(port string) *Server {
	return &Server{port: port}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var root string
	root, r.URL.Path = shiftPath(r.URL.Path)

	enableCors(&w)

	switch root {
	case "debug":
		s.handleDebug(w, r)
	case "identity":
		s.handleIdentity(w, r)
	case "mocks":
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

// shiftPath splits off the first fragment of p.
// head will never contain a slash and tail will
// always be a rooted path without trailing slash.
// 	/foo/bar?x=1234 -> foo, /bar?x=1234
func shiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}
