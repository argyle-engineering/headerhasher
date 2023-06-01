package headerhasher_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/argyle-engineering/headerhasher"
)

func TestDemo(t *testing.T) {
	cfg := headerhasher.CreateConfig()

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := headerhasher.New(ctx, next, cfg, "test-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	req.Header.Set("Authorization", "hello")
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	assertHeader(t, req, "Authorization", "hello")
	assertHeader(t, req, "Authorization-Hashed", "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824")
}

func assertHeader(t *testing.T, req *http.Request, key, expected string) {
	t.Helper()

	if req.Header.Get(key) != expected {
		t.Errorf("invalid header value: %s", req.Header.Get(key))
	}
}
