package main
import (
	"testing"
	"net/http/httptest"
	"net/http"
)

func TestServeHTTP(t *testing.T) {
	var h *authHandler
	ts := httptest.NewServer( http.HandleFunc( h.next.ServeHTTP ) )

	defer ts.Close()

	res, err := http.Get( ts.URL )
	if err != nil {
		t.Error("unexpected")
		return
	}

	if res.StatusCode != 200 {
		t.Error("Status code error")
		return
	}
}