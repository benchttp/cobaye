package server

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type paramkey string

const (
	paramkeyDelay paramkey = "delay"
	paramkeyFib   paramkey = "fib"
)

func (s *Server) handleFibonacci(w http.ResponseWriter, r *http.Request) {
	s.incrementRequestCount()

	params := r.URL.Query()

	delay, err := readParamDuration(params, paramkeyDelay)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	fibInt, err := readParamInt(params, paramkeyFib)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	if delay > 0 {
		time.Sleep(delay)
	}

	if fibInt > 0 {
		_ = fibonacci(fibInt)
	}
}

func readParamInt(params url.Values, key paramkey) (int, error) {
	raw := params.Get(string(key))
	if raw == "" {
		return 0, nil
	}

	n, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("invalid param: %s: want int, got %s", key, raw)
	}

	return n, nil
}

func readParamDuration(params url.Values, key paramkey) (time.Duration, error) {
	raw := params.Get(string(key))
	if raw == "" {
		return 0, nil
	}

	d, err := time.ParseDuration(raw)
	if err != nil {
		return 0, fmt.Errorf("invalid param: %s: want parsable duration, got %s", key, raw)
	}

	return d, nil
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
