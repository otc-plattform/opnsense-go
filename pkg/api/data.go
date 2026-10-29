package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strconv"
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
	BaseEndpoint    string
	Method          string
	PathParameters  []string
	QueryParameters map[string]string
	BodyParameters  map[string]interface{}
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

	if len(p.QueryParameters) > 0 {
		keys := make([]string, 0, len(p.QueryParameters))
		for k := range p.QueryParameters {
			keys = append(keys, k)
		}
		// Sort so the URL is deterministic — important for any test that
		// asserts on the request URL and for caches/loggers that key on it.
		sort.Strings(keys)

		values := url.Values{}
		for _, k := range keys {
			values.Set(k, p.QueryParameters[k])
		}

		separator := "?"
		if strings.Contains(currentPath, "?") {
			separator = "&"
		}
		currentPath += separator + values.Encode()
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
	Status      string            `json:"status"`
	Result      string            `json:"result"`
	UUID        string            `json:"uuid"`
	Validations map[string]string `json:"validations"`
}

type Option struct {
	Value    string `json:"value"`
	Selected int    `json:"selected"`
}

type FieldOptions map[string]Option

type BoolLike bool

func (b *BoolLike) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "true", "false":
		var v bool
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		*b = BoolLike(v)
		return nil
	}

	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		switch strings.ToLower(strings.TrimSpace(str)) {
		case "1", "true", "yes", "on":
			*b = BoolLike(true)
			return nil
		case "0", "false", "no", "off", "":
			*b = BoolLike(false)
			return nil
		}
	}

	var n int
	if err := json.Unmarshal(data, &n); err == nil {
		*b = BoolLike(n != 0)
		return nil
	}

	return fmt.Errorf("BoolLike: unsupported value %s", string(data))
}

func (b BoolLike) Bool() bool { return bool(b) }

// IntLike accepts JSON integers encoded as numbers or strings. Empty strings
// and null decode to zero, matching the OPNsense API's optional fields.
type IntLike int

func (i *IntLike) UnmarshalJSON(data []byte) error {
	var num json.Number
	if err := json.Unmarshal(data, &num); err == nil && num.String() != "" {
		v, err := strconv.Atoi(num.String())
		if err != nil {
			return fmt.Errorf("IntLike: invalid numeric value %q: %w", num.String(), err)
		}
		*i = IntLike(v)
		return nil
	}

	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		s = strings.TrimSpace(s)
		if s == "" {
			*i = 0
			return nil
		}
		v, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("IntLike: invalid string %q: %w", s, err)
		}
		*i = IntLike(v)
		return nil
	}

	if string(data) == "null" {
		*i = 0
		return nil
	}

	return fmt.Errorf("IntLike: unsupported JSON token %s", string(data))
}

func (i IntLike) Int() int { return int(i) }

func (o *FieldOptions) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if bytes.Equal(data, []byte("null")) || len(data) == 0 {
		*o = nil
		return nil
	}

	switch data[0] {
	case '{':
		var m map[string]Option
		if err := json.Unmarshal(data, &m); err != nil {
			return err
		}
		*o = m
		return nil
	case '[':
		if bytes.Equal(data, []byte("[]")) {
			*o = FieldOptions{}
			return nil
		}
		return fmt.Errorf("FieldOptions: expected object or empty array, got %s", data)
	default:
		return fmt.Errorf("FieldOptions: invalid JSON starting with %q", data[0])
	}
}

type ReconfigureStatusResult struct {
	Status string `json:"status"`
}
