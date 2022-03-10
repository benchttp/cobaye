package server

import (
	"fmt"
	"io"
	"net/http"
)

func (s *Server) handleIdentity(w http.ResponseWriter, r *http.Request) {
	for key, vals := range r.Header {
		for _, val := range vals {
			w.Header().Add(key, val)
		}
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(b)
}
