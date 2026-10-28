package haproxy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/errs"
)

func (c *Controller) HAProxySearchHealthMonitors(ctx context.Context) (*HAProxyObjectSearchResult, error) {
	return c.searchObjects(ctx, "/haproxy/settings/search_healthchecks", "SearchHealthMonitors")
}

func (c *Controller) HAProxyGetHealthMonitor(ctx context.Context, uuid string) (HAProxyObject, error) {
	opts := api.RPCOpts{BaseEndpoint: "/haproxy/settings/get_healthcheck", Method: "GET", PathParameters: []string{uuid}}
	var raw json.RawMessage
	result, err := api.Call(c.Client(), ctx, opts, &raw)
	if err != nil {
		return nil, fmt.Errorf("GetHealthMonitor call failed: %w", err)
	}
	if bytes.Equal(bytes.TrimSpace(*result), []byte("[]")) || bytes.Equal(bytes.TrimSpace(*result), []byte("null")) {
		return nil, errs.NewNotFoundError()
	}
	var response struct {
		Healthcheck HAProxyObject `json:"healthcheck"`
	}
	if err := json.Unmarshal(*result, &response); err != nil {
		return nil, fmt.Errorf("GetHealthMonitor response: %w", err)
	}
	if len(response.Healthcheck) == 0 {
		return nil, errs.NewNotFoundError()
	}
	return response.Healthcheck, nil
}

func (c *Controller) HAProxyAddHealthMonitor(ctx context.Context, monitor HAProxyObject) (*api.ActionResult, error) {
	return c.addObject(ctx, "/haproxy/settings/add_healthcheck", "healthcheck", "AddHealthMonitor", monitor)
}

func (c *Controller) HAProxyEditHealthMonitor(ctx context.Context, uuid string, monitor HAProxyObject) (*api.ActionResult, error) {
	return c.editObject(ctx, "/haproxy/settings/set_healthcheck", "healthcheck", "EditHealthMonitor", uuid, monitor)
}

func (c *Controller) HAProxyDeleteHealthMonitor(ctx context.Context, uuid string) (*api.ActionResult, error) {
	return c.deleteObject(ctx, "/haproxy/settings/del_healthcheck", "DeleteHealthMonitor", uuid)
}
