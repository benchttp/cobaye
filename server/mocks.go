package server

import (
	"fmt"
	"net/http"
)

func mockFile(file string) string {
	return fmt.Sprintf("./mocks/%s", file)
}

func (s *Server) handleMocks(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.NotFound(w, r)
		return
	}

	var root string
	root, r.URL.Path = shiftPath(r.URL.Path)

	switch root {
	case "reports":
		if r.URL.Path == "/" {
			// There is no GET /reports
			http.NotFound(w, r)
		} else {
			// Serve mock of /reports/{id}
			http.ServeFile(w, r, mockFile("report.json"))
		}

	case "stats":
		if r.URL.Path == "/" {
			// Serve mock of /stats
			http.ServeFile(w, r, mockFile("stats-list.json"))
		} else {
			// Serve mock of /stats/{id}
			http.ServeFile(w, r, mockFile("stats.json"))
		}

	default:
		http.NotFound(w, r)
	}
}
