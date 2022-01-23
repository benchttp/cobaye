package main //nolint:testpackage

import (
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/drykit-go/testx"
	"github.com/drykit-go/testx/check"
)

func TestHandleRequest(t *testing.T) {
	t.Run("request with delay param", func(t *testing.T) {
		const (
			delay  = 100 * time.Millisecond
			margin = 5 * time.Millisecond
			expmin = delay
			expmax = delay + margin
		)

		r := httptest.NewRequest("", fmt.Sprintf("/?delay=%dms", delay.Milliseconds()), nil)

		testx.HTTPHandlerFunc(handleRequest).WithRequest(r).
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

		testx.HTTPHandlerFunc(handleRequest).WithRequest(r).
			Response(checkStatusCode(200)).
			Duration(check.Duration.InRange(expmin, expmax)).
			Run(t)
	})

	t.Run("request without params", func(t *testing.T) {
		const expmax = 3 * time.Millisecond

		testx.HTTPHandlerFunc(handleRequest).
			Response(checkStatusCode(200)).
			Duration(check.Duration.Under(expmax)).
			Run(t)
	})

	t.Run("request with invalid params", func(t *testing.T) {
		const expmax = 3 * time.Millisecond

		r := httptest.NewRequest("", "/?delay=hey&fib=100", nil)

		testx.HTTPHandlerFunc(handleRequest).WithRequest(r).
			Response(checkStatusCode(400)).
			Duration(check.Duration.Under(expmax)).
			Run(t)
	})
}

func checkStatusCode(code int) check.HTTPResponseChecker {
	return check.HTTPResponse.StatusCode(check.Int.Is(code))
}
