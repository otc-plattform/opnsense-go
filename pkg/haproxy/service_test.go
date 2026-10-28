package haproxy

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/browningluke/opnsense-go/pkg/api"
)

func TestHAProxyConfigTestResponse(t *testing.T) {
	for _, tt := range []struct {
		name, body, want string
		wantError        bool
	}{
		{"empty result", `{"result":""}`, "", false},
		{"output", `{"result":"Configuration file is valid"}`, "Configuration file is valid", false},
		{"missing result", `{}`, "", true},
		{"null result", `{"result":null}`, "", true},
		{"invalid type", `{"result":{}}`, "", true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodPost || r.URL.Path != "/api/haproxy/service/configtest" {
					t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
				}
				io.WriteString(w, tt.body)
			}))
			defer server.Close()
			c := Controller{Api: api.NewClient(api.Options{Uri: server.URL, Logger: log.New(io.Discard, "", 0)})}
			got, err := c.HAProxyConfigTest(context.Background())
			if (err != nil) != tt.wantError || (!tt.wantError && got != tt.want) {
				t.Fatalf("got %q, %v", got, err)
			}
		})
	}
}
