package server_test

import (
	"fmt"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/drykit-go/testx"
	"github.com/drykit-go/testx/check"

	"github.com/benchttp/cobaye/internal/server"
)

func TestHandleRequest(t *testing.T) {
	s := &server.Server{}

	t.Run("request with delay param", func(t *testing.T) {
		const (
			delay  = 100 * time.Millisecond
			margin = 5 * time.Millisecond
			expmin = delay
			expmax = delay + margin
		)

		r := httptest.NewRequest("", fmt.Sprintf("/?delay=%dms", delay.Milliseconds()), nil)

		testx.HTTPHandler(s).WithRequest(r).
			Response(checkStatusCode(200)).
			Duration(check.Duration.InRange(expmin, expmax)).
			Run(t)
	})

	t.Run("request without params", func(t *testing.T) {
		const expmax = 3 * time.Millisecond

		testx.HTTPHandler(s).
			Response(checkStatusCode(200)).
			Duration(check.Duration.Under(expmax)).
			Run(t)
	})

	t.Run("request with invalid params", func(t *testing.T) {
		const expmax = 3 * time.Millisecond

		r := httptest.NewRequest("", "/?delay=hey&fib=100", nil)

		testx.HTTPHandler(s).WithRequest(r).
			Response(checkStatusCode(400)).
			Duration(check.Duration.Under(expmax)).
			Run(t)
	})
}

func TestHandleDebug(t *testing.T) {
	t.Run("initialized with 0 request", func(t *testing.T) {
		const expRequests = 0

		s := &server.Server{}
		r := httptest.NewRequest("", "/debug", nil)
		testx.HTTPHandler(s).WithRequest(r).
			Response(
				checkStatusCode(200),
				checkExactBody([]byte([]byte(strconv.Itoa(expRequests)))),
			).
			Run(t)
	})

	t.Run("count requests", func(t *testing.T) {
		const expRequests = 42

		s := &server.Server{}
		regularRequest := httptest.NewRequest("", "/", nil)

		for i := 0; i < expRequests; i++ {
			s.ServeHTTP(nil, regularRequest)
		}

		debugRequest := httptest.NewRequest("", "/debug", nil)
		testx.HTTPHandler(s).WithRequest(debugRequest).
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
