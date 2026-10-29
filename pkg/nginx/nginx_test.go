package nginx

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGetResponsesDecodeOptionMapsAsStringLists(t *testing.T) {
	var location LocationGetResponse
	if err := json.Unmarshal([]byte(`{"location":{"index":{"index.html":{"value":"index.html","selected":1}}}}`), &location); err != nil {
		t.Fatalf("unmarshal location response: %v", err)
	}
	if got, want := []string(location.Location.Index), []string{"index.html"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("location index = %#v, want %#v", got, want)
	}

	var server HTTPServerGetResponse
	if err := json.Unmarshal([]byte(`{"httpserver":{"trusted_proxies":{"192.0.2.1":{"value":"192.0.2.1","selected":1}}}}`), &server); err != nil {
		t.Fatalf("unmarshal HTTP server response: %v", err)
	}
	if got, want := []string(server.HTTPServer.TrustedProxies), []string{"192.0.2.1"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("trusted proxies = %#v, want %#v", got, want)
	}
}
