package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type ReqOpts struct {
	AddEndpoint         string
	GetEndpoint         string
	UpdateEndpoint      string
	DeleteEndpoint      string
	ReconfigureEndpoint string

	Monad string
}

// Response structs
type addResp struct {
	Result      string                 `json:"result"`
	UUID        string                 `json:"uuid"`
	Validations map[string]interface{} `json:"validations,omitempty"`
}

type deleteResp struct {
	Result string `json:"result"`
}

// RCP Options
type RPCOpts struct {
	BaseEndpoint   string
	Method         string
	PathParameters []string
	BodyParameters map[string]interface{}
}

func (p *RPCOpts) EndpointURL() string {
	currentPath := p.BaseEndpoint
	for _, param := range p.PathParameters {
		escapedParam := url.PathEscape(param)

		if currentPath == "" {
			currentPath = escapedParam
		} else if strings.HasSuffix(currentPath, "/") {
			currentPath += escapedParam
		} else {
			currentPath += "/" + escapedParam
		}
	}
	return currentPath
}

func (p *RPCOpts) Body() (string, error) {
	if len(p.BodyParameters) == 0 {
		return "{}", nil
	}
	jsonBytes, err := json.Marshal(p.BodyParameters)
	if err != nil {
		return "", fmt.Errorf("failed to marshal BodyParameters to JSON: %w", err)
	}
	return string(jsonBytes), nil
}

type ActionResult struct {
	Result      string            `json:"result"`
	UUID        string            `json:"uuid"`
	Validations map[string]string `json:"validations"`
}

type Option struct {
	Value    string `json:"value"`
	Selected int    `json:"selected"`
}

type FieldOptions map[string]Option

func (o *FieldOptions) UnmarshalJSON(b []byte) error {
	b = bytes.TrimSpace(b)

	if bytes.Equal(b, []byte("null")) || len(b) == 0 {
		*o = nil
		return nil
	}

	switch b[0] {
	case '{':
		var m map[string]Option
		if err := json.Unmarshal(b, &m); err != nil {
			return err
		}
		*o = m
		return nil

	case '[':
		// Accept exactly the empty array as "empty map"
		if bytes.Equal(b, []byte("[]")) {
			*o = FieldOptions{}
			return nil
		}
		return fmt.Errorf("Options: expected object or empty array, got %s", b)

	default:
		return fmt.Errorf("Options: invalid JSON starting with %q", b[0])
	}
}
