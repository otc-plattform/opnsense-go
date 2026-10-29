package haproxy

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/errs"
)

func TestGetResolver(t *testing.T) {
	for _, tc := range []struct {
		body     string
		notFound bool
	}{
		{`{"resolver":{"name":"dns"}}`, false}, {`[]`, true}, {`{}`, true},
	} {
		t.Run(tc.body, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "GET" || r.URL.Path != "/api/haproxy/settings/get_resolver/test" {
					t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
				}
				io.WriteString(w, tc.body)
			}))
			defer server.Close()
			c := Controller{Api: api.NewClient(api.Options{Uri: server.URL, Logger: log.New(io.Discard, "", 0)})}
			obj, err := c.HAProxyGetResolver(context.Background(), "test")
			var notFound *errs.NotFoundError
			if tc.notFound {
				if !errors.As(err, &notFound) {
					t.Fatalf("expected not found, got %v", err)
				}
				return
			}
			if err != nil || obj["name"] != "dns" {
				t.Fatalf("got %v, %v", obj, err)
			}
		})
	}
}
