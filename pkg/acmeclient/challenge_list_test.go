package acmeclient

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/browningluke/opnsense-go/pkg/api"
)

func TestGetChallengeTypeListFormats(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/api/acmeclient/validations/get/test-id" {
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			http.Error(w, "unexpected request", http.StatusNotFound)
			return
		}
		io.WriteString(w, `{"validation":{"name":"test","http_opn_ipaddresses":{},"http_haproxyFrontends":["frontend-1"],"tlsalpn_acme_ipaddresses":{"":{"value":"","selected":1}}}}`)
	}))
	defer server.Close()

	controller := Controller{Api: api.NewClient(api.Options{Uri: server.URL, Logger: log.New(io.Discard, "", 0)})}
	result, err := controller.ACMEClientGetChallengeType(context.Background(), "test-id")
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Validation.HTTPOpnIPAddresses) != 0 || len(result.Validation.TLSALPNAcmeIPAddresses) != 0 {
		t.Fatalf("expected empty address lists, got %#v and %#v", result.Validation.HTTPOpnIPAddresses, result.Validation.TLSALPNAcmeIPAddresses)
	}
	if !reflect.DeepEqual(result.Validation.HTTPHAProxyFrontends, api.StringList{"frontend-1"}) {
		t.Fatalf("unexpected frontends: %#v", result.Validation.HTTPHAProxyFrontends)
	}
}
