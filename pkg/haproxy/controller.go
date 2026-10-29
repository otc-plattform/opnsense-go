package haproxy

import "github.com/browningluke/opnsense-go/pkg/api"

// Controller for HAProxy.
type Controller struct {
	Api *api.Client
}

func (c *Controller) Client() *api.Client {
	return c.Api
}
