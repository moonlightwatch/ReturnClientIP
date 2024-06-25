package ReturnClientIP_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/moonlightwatch/ReturnClientIP"
)

func TestReturnClientIP(t *testing.T) {
	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})
	cfg := ReturnClientIP.CreateConfig()

	rcIP, err := ReturnClientIP.New(ctx, next, cfg, "test")
	if err != nil {
		t.Fatalf("failed to create new handler: %v", err)
	}
	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.RemoteAddr = "127.0.0.1:1234"
	rcIP.ServeHTTP(recorder, req)

	if recorder.Body.String() != "127.0.0.1" {
		t.Errorf("Expected 127.0.0.1, got %s", recorder.Body.String())
	}

}
