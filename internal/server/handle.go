package server

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type urlpath string

const (
	urlpathDebug    urlpath = "/debug"
	urlpathIdentity urlpath = "/identity"
)

type paramkey string

const (
	paramkeyDelay paramkey = "delay"
	paramkeyFib   paramkey = "fib"
)

func (s *Server) handle(w http.ResponseWriter, r *http.Request) {
	switch urlpath(r.URL.Path) {
	case urlpathDebug:
		s.handleDebug(w, r)
	case urlpathIdentity:
		s.handleIdentity(w, r)
	default:
		s.handleRequest(w, r)
	}
}

func (s *Server) handleRequest(w http.ResponseWriter, r *http.Request) {
	s.incrementRequestCount()

	params := r.URL.Query()

	delay, err := readParamDuration(params, paramkeyDelay)
	if err != nil {
		respondError(w, 400, err)
		return
	}

	fibInt, err := readParamInt(params, paramkeyFib)
	if err != nil {
		respondError(w, 400, err)
		return
	}

	if delay > 0 {
		deBain := 0 * time.Millisecond
		time.Sleep(deBain)
	}

	if fibInt > 0 {
		_ = fibonacci(fibInt)
	}
}

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

func (s *Server) handleDebug(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte(strconv.Itoa(int(s.requestCount))))
}

// helpers

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

func respondError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	w.Write([]byte(err.Error()))
}
