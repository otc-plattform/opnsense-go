package haproxy

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/errs"
)

func TestHealthMonitorOperations(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method + " " + r.URL.Path {
		case "POST /api/haproxy/settings/add_healthcheck", "POST /api/haproxy/settings/set_healthcheck/id":
			var body struct {
				Healthcheck HAProxyObject `json:"healthcheck"`
			}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.Healthcheck["name"] != "test" {
				t.Errorf("invalid healthcheck request: %v, %#v", err, body)
			}
			io.WriteString(w, `{"result":"saved","uuid":"id"}`)
		case "GET /api/haproxy/settings/get_healthcheck/id":
			io.WriteString(w, `{"healthcheck":{"name":"test"}}`)
		case "GET /api/haproxy/settings/get_healthcheck/missing":
			io.WriteString(w, `[]`)
		case "POST /api/haproxy/settings/search_healthchecks":
			io.WriteString(w, `{"rows":[{"uuid":"id","name":"test"}]}`)
		case "POST /api/haproxy/settings/del_healthcheck/id":
			io.WriteString(w, `{"result":"deleted"}`)
		default:
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
			http.NotFound(w, r)
		}
	}))
	defer server.Close()
	c := Controller{Api: api.NewClient(api.Options{Uri: server.URL, Logger: log.New(io.Discard, "", 0)})}
	ctx := context.Background()
	if result, err := c.HAProxyAddHealthMonitor(ctx, HAProxyObject{"name": "test"}); err != nil || result.UUID != "id" {
		t.Fatalf("add: %v %v", result, err)
	}
	if result, err := c.HAProxyGetHealthMonitor(ctx, "id"); err != nil || result["name"] != "test" {
		t.Fatalf("get: %v %v", result, err)
	}
	if result, err := c.HAProxyEditHealthMonitor(ctx, "id", HAProxyObject{"name": "test"}); err != nil || result.Result != "saved" {
		t.Fatalf("edit: %v %v", result, err)
	}
	if result, err := c.HAProxySearchHealthMonitors(ctx); err != nil || len(result.Rows) != 1 {
		t.Fatalf("search: %v %v", result, err)
	}
	if result, err := c.HAProxyDeleteHealthMonitor(ctx, "id"); err != nil || result.Result != "deleted" {
		t.Fatalf("delete: %v %v", result, err)
	}
	_, err := c.HAProxyGetHealthMonitor(ctx, "missing")
	var notFound *errs.NotFoundError
	if !errors.As(err, &notFound) {
		t.Fatalf("missing: expected not found, got %v", err)
	}
}
