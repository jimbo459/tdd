package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := makeDelayedWebServer(time.Second * 5)
	fastServer := makeDelayedWebServer(time.Second * 1)

	defer slowServer.Close()
	defer fastServer.Close()

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	want := fastUrl

	got, _ := Racer(slowUrl, fastUrl)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	t.Run("returns an error if Server does not respond within 10s", func(t *testing.T) {
		serverA := makeDelayedWebServer(time.Second * 11)
		serverB := makeDelayedWebServer(time.Second * 12)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)

		if err == nil {
			t.Errorf("Expected error!")
		}
	})
}

func makeDelayedWebServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
