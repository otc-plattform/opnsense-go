package api

import (
	"net/http"
	"testing"
)

func TestNewClientDisablesCompression(t *testing.T) {
	client := NewClient(Options{})

	transport, ok := client.client.HTTPClient.Transport.(*http.Transport)
	if !ok {
		t.Fatalf("transport type = %T, want *http.Transport", client.client.HTTPClient.Transport)
	}
	if !transport.DisableCompression {
		t.Fatal("DisableCompression = false, want true")
	}
}
