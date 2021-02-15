// +build go1.9

package mux

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSchemeMatchers(t *testing.T) {
	router := NewRouter()
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("main http world"))
	}).Schemes("http")
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("main https world"))
	}).Schemes("https")

	assertResponseBody := func(t *testing.T, s *httptest.Server, expectedBody string) {
		resp, err := s.Client().Get(s.URL)
		if err != nil {
			t.Fatalf("unexpected error getting from server: %v", err)
		}
		if resp.StatusCode != 200 {
			t.Fatalf("expected a status code of 200, got %v", resp.StatusCode)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("unexpected error reading body: %v", err)
		}
		if !bytes.Equal(body, []byte(expectedBody)) {
			t.Fatalf("response should be main world, was: %q", string(body))
		}
	}

	t.Run("httpServer", func(t *testing.T) {
		s := httptest.NewServer(router)
		defer s.Close()
		assertResponseBody(t, s, "main http world")
	})
	t.Run("httpsServer", func(t *testing.T) {
		s := httptest.NewTLSServer(router)
		defer s.Close()
		assertResponseBody(t, s, "main https world")
	})
}
