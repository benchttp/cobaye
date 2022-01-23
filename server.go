package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

type server struct {
	mu            sync.Mutex
	totalRequests int
}

func (s *server) listenInput() error {
	reader := bufio.NewReader(os.Stdin)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("yo")
			return err
		}

		if string(line) == "debug" {
			fmt.Printf("Total requests: %d\n", s.totalRequests)
		}
	}
}

func (s *server) incrementTotalRequests() {
	s.mu.Lock()
	s.totalRequests++
	s.mu.Unlock()
}
