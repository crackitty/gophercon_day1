package routing_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/crackitty/gophercon_day1/cmd/routing"
)

func TestBaseRouter(t *testing.T) {
	handler := routing.BaseRouter()
	ts := httptest.NewServer(handler)
	res, err := http.Get(ts.URL + "/bubbles")

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fail()
	}
}
