package main //nolint:testpackage

import (
	"fmt"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/drykit-go/testx"
	"github.com/drykit-go/testx/check"
)

func TestHandleRequest(t *testing.T) {
	s := server{}

	t.Run("request with delay param", func(t *testing.T) {
		const (
			delay  = 100 * time.Millisecond
			margin = 5 * time.Millisecond
			expmin = delay
			expmax = delay + margin
		)

		r := httptest.NewRequest("", fmt.Sprintf("/?delay=%dms", delay.Milliseconds()), nil)

		testx.HTTPHandlerFunc(s.handleRequest).WithRequest(r).
			Response(checkStatusCode(200)).
			Duration(check.Duration.InRange(expmin, expmax)).
			Run(t)
	})

	t.Run("request with fib param", func(t *testing.T) {
		const (
			fib    = 35
			expmin = 30 * time.Millisecond
			expmax = 80 * time.Millisecond
		)

		r := httptest.NewRequest("", fmt.Sprintf("/?fib=%d", fib), nil)

		testx.HTTPHandlerFunc(s.handleRequest).WithRequest(r).
			Response(checkStatusCode(200)).
			Duration(check.Duration.InRange(expmin, expmax)).
			Run(t)
	})

	t.Run("request without params", func(t *testing.T) {
		const expmax = 3 * time.Millisecond

		testx.HTTPHandlerFunc(s.handleRequest).
			Response(checkStatusCode(200)).
			Duration(check.Duration.Under(expmax)).
			Run(t)
	})

	t.Run("request with invalid params", func(t *testing.T) {
		const expmax = 3 * time.Millisecond

		r := httptest.NewRequest("", "/?delay=hey&fib=100", nil)

		testx.HTTPHandlerFunc(s.handleRequest).WithRequest(r).
			Response(checkStatusCode(400)).
			Duration(check.Duration.Under(expmax)).
			Run(t)
	})
}

func TestHandleDebug(t *testing.T) {
	t.Run("initialized with 0 request", func(t *testing.T) {
		s := server{}
		r := httptest.NewRequest("", "/debug", nil)
		testx.HTTPHandlerFunc(s.handleDebug).WithRequest(r).
			Response(
				checkStatusCode(200),
				checkExactBody([]byte("0")),
			).
			Run(t)
	})

	t.Run("count requests", func(t *testing.T) {
		const expRequests = 42
		s := server{}
		regularRequest := httptest.NewRequest("", "/", nil)

		for i := 0; i < expRequests; i++ {
			s.handleRequest(nil, regularRequest)
		}

		debugRequest := httptest.NewRequest("", "/debug", nil)
		testx.HTTPHandlerFunc(s.handleDebug).WithRequest(debugRequest).
			Response(
				checkStatusCode(200),
				checkExactBody([]byte(strconv.Itoa(expRequests))),
			).
			Run(t)
	})
}

// helpers

func checkStatusCode(code int) check.HTTPResponseChecker {
	return check.HTTPResponse.StatusCode(check.Int.Is(code))
}

func checkExactBody(value []byte) check.HTTPResponseChecker {
	return check.HTTPResponse.Body(check.Bytes.Is(value))
}
