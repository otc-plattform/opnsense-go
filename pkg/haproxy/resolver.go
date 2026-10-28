package haproxy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/errs"
)

func (c *Controller) HAProxySearchResolvers(ctx context.Context) (*HAProxyObjectSearchResult, error) {
	return c.searchObjects(ctx, "/haproxy/settings/search_resolvers", "SearchResolvers")
}

func (c *Controller) HAProxyGetResolver(ctx context.Context, uuid string) (HAProxyObject, error) {
	opts := api.RPCOpts{BaseEndpoint: "/haproxy/settings/get_resolver", Method: "GET", PathParameters: []string{uuid}}
	var raw json.RawMessage
	result, err := api.Call(c.Client(), ctx, opts, &raw)
	if err != nil {
		return nil, fmt.Errorf("GetResolver call failed: %w", err)
	}
	if bytes.Equal(bytes.TrimSpace(*result), []byte("[]")) || bytes.Equal(bytes.TrimSpace(*result), []byte("null")) {
		return nil, errs.NewNotFoundError()
	}
	var response struct {
		Resolver HAProxyObject `json:"resolver"`
	}
	if err := json.Unmarshal(*result, &response); err != nil {
		return nil, fmt.Errorf("GetResolver response: %w", err)
	}
	if len(response.Resolver) == 0 {
		return nil, errs.NewNotFoundError()
	}
	return response.Resolver, nil
}

func (c *Controller) HAProxyAddResolver(ctx context.Context, resolver HAProxyObject) (*api.ActionResult, error) {
	return c.addObject(ctx, "/haproxy/settings/add_resolver", "resolver", "AddResolver", resolver)
}

func (c *Controller) HAProxyEditResolver(ctx context.Context, uuid string, resolver HAProxyObject) (*api.ActionResult, error) {
	return c.editObject(ctx, "/haproxy/settings/set_resolver", "resolver", "EditResolver", uuid, resolver)
}

func (c *Controller) HAProxyDeleteResolver(ctx context.Context, uuid string) (*api.ActionResult, error) {
	return c.deleteObject(ctx, "/haproxy/settings/del_resolver", "DeleteResolver", uuid)
}
