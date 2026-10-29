package firmware

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/browningluke/opnsense-go/pkg/api"
)

func TestFirmwareInfoProductCheck(t *testing.T) {
	tests := []struct {
		name  string
		value string
		check func(t *testing.T, value any)
	}{
		{"null", `null`, func(t *testing.T, value any) {
			if value != nil {
				t.Fatalf("got %#v, want nil", value)
			}
		}},
		{"string", `"ok"`, func(t *testing.T, value any) {
			if value != "ok" {
				t.Fatalf("got %#v, want ok", value)
			}
		}},
		{"object", `{"status":"ok","upgrade_packages":[]}`, func(t *testing.T, value any) {
			object, ok := value.(map[string]any)
			if !ok || object["status"] != "ok" {
				t.Fatalf("got %#v, want object with status ok", value)
			}
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet || r.URL.Path != "/api/core/firmware/info" {
					t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
					http.Error(w, "unexpected request", http.StatusNotFound)
					return
				}
				fmt.Fprintf(w, `{"product":{"product_check":%s},"package":[{"name":"os-acme-client","installed":"1"}]}`, tt.value)
			}))
			defer server.Close()

			controller := Controller{Api: api.NewClient(api.Options{Uri: server.URL, Logger: log.New(io.Discard, "", 0)})}
			result, err := controller.FirmwareInfo(context.Background())
			if err != nil {
				t.Fatalf("FirmwareInfo: %v", err)
			}
			tt.check(t, result.Product.ProductCheck)
			if len(result.Package) != 1 || result.Package[0].Name != "os-acme-client" || result.Package[0].Installed != "1" {
				t.Fatalf("unexpected packages: %#v", result.Package)
			}
		})
	}
}
