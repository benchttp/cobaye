package server

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
)

func (s *Server) handleDebug(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte(strconv.Itoa(int(s.requestCount))))
}

func (s *Server) ListenStdin() {
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
