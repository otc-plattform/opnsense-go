package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
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
	// trim quotes if present
	s := string(data)
	switch s {
	case "true", "false":
		var v bool
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		*b = BoolLike(v)
		return nil
	}

	// handle strings: "1","0","true","false",""
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

	// handle numbers 0/1 (rare)
	var n int
	if err := json.Unmarshal(data, &n); err == nil {
		*b = BoolLike(n != 0)
		return nil
	}

	return fmt.Errorf("BoolLike: unsupported value %s", string(data))
}

// Bool returns the concrete bool value.
func (b BoolLike) Bool() bool { return bool(b) }

// IntLike handles JSON that may encode an integer as either a number or a string.
// Examples it will accept: 254, "254", "", "  254  ".
// Empty string and null decode to 0.
type IntLike int

// UnmarshalJSON implements json.Unmarshaler for IntLike.
func (i *IntLike) UnmarshalJSON(data []byte) error {
	// Fast path: genuine JSON number
	// We try json.Number first to avoid float rounding.
	var num json.Number
	if err := json.Unmarshal(data, &num); err == nil {
		if numStr := num.String(); numStr != "" {
			v, err := strconv.Atoi(numStr)
			if err != nil {
				return fmt.Errorf("IntLike: invalid numeric value %q: %w", numStr, err)
			}
			*i = IntLike(v)
			return nil
		}
	}

	// Handle string-encoded numbers (or empty)
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

	// Treat explicit null as zero
	if string(data) == "null" {
		*i = 0
		return nil
	}

	return fmt.Errorf("IntLike: unsupported JSON token %s", string(data))
}

// Int returns the concrete int value.
func (i IntLike) Int() int { return int(i) }

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
