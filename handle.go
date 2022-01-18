package main

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

func handle(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	delayMS, err := readParamInt(params, paramkeyDelay)
	if err != nil {
		respondError(w, 400, err)
		return
	}

	fibInt, err := readParamInt(params, paramkeyFib)
	if err != nil {
		respondError(w, 400, err)
		return
	}

	if delayMS > 0 {
		time.Sleep(time.Duration(delayMS) * time.Millisecond)
	}

	if fibInt > 0 {
		fibonacci(fibInt) //nolint
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

func respondError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	w.Write([]byte(err.Error()))
}
