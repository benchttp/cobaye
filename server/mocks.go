package server

import (
	"fmt"
	"net/http"
)

func mockFile(file string) string {
	return fmt.Sprintf("./mocks/%s", file)
}

func (s *Server) handleMocks(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, mockFile("report-get.json"))
}
