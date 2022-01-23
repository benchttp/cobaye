package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"
)

type urlpath string

const urlpathDebug urlpath = "/debug"

type paramkey string

const (
	paramkeyDelay paramkey = "delay"
	paramkeyFib   paramkey = "fib"
)

var (
	mu            = sync.Mutex{}
	totalRequests = 0
)

func handleMain(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case string(urlpathDebug):
		handleDebug(w, r)
	default:
		handleRequest(w, r)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	incrementTotalRequests()

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
		time.Sleep(delay)
	}

	if fibInt > 0 {
		fibonacci(fibInt) //nolint
	}
}

func handleDebug(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte(strconv.Itoa(totalRequests)))
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

func respondError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	w.Write([]byte(err.Error()))
}

func incrementTotalRequests() {
	mu.Lock()
	totalRequests++
	mu.Unlock()
}
