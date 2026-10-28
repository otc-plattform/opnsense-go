package haproxy

import (
	"context"
	"fmt"
	"strings"

	"github.com/browningluke/opnsense-go/pkg/api"
)

// HAProxyConfigTest regenerates and checks the staged HAProxy configuration.
// OPNsense returns the HAProxy output as a string even if the test fails.
func (c *Controller) HAProxyConfigTest(ctx context.Context) (string, error) {
	opts := api.RPCOpts{BaseEndpoint: "/haproxy/service/configtest", Method: "POST"}
	var response struct {
		Result *string `json:"result"`
	}
	result, err := api.Call(c.Client(), ctx, opts, &response)
	if err != nil {
		return "", fmt.Errorf("HAProxy config test: %w", err)
	}
	if result.Result == nil {
		return "", fmt.Errorf("HAProxy config test returned no result")
	}
	return *result.Result, nil
}

// HAProxyReconfigure activates the staged HAProxy configuration.
func (c *Controller) HAProxyReconfigure(ctx context.Context) error {
	opts := api.RPCOpts{BaseEndpoint: "/haproxy/service/reconfigure", Method: "POST"}
	var response struct {
		Status string `json:"status"`
		Result string `json:"result"`
	}
	result, err := api.Call(c.Client(), ctx, opts, &response)
	if err != nil {
		return fmt.Errorf("HAProxy reconfigure: %w", err)
	}
	status := result.Status
	if status == "" {
		status = result.Result
	}
	if !strings.EqualFold(strings.TrimSpace(status), "ok") {
		return fmt.Errorf("HAProxy reconfigure returned status %q", status)
	}
	return nil
}
