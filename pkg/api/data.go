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

// StringList handles JSON that may encode a list of strings in any of the
// shapes OPNsense uses for Multiple/list fields, depending on field type and
// API version (see the acmeclient validations GET):
//
//	null / missing           -> empty list
//	["a","b"]                -> ["a","b"]                  (plain array)
//	{"a":"a","b":"b"}        -> ["a","b"]                  (MultiSelect: key == value)
//	{"<uuid>":"label"}       -> ["<uuid>"]                 (relation: key == UUID)
//	{"0":"a","1":"b"}        -> ["a","b"]                  (index -> value)
//	"a,b"                    -> ["a","b"]                  (raw CSV string)
//	"a"                      -> ["a"]                      (single entry)
//
// OPNsense otherwise serializes some of these fields as JSON objects, which
// breaks strict []string decoding.
type StringList []string

// UnmarshalJSON implements json.Unmarshaler for StringList.
func (l *StringList) UnmarshalJSON(data []byte) error {
	trimmed := bytes.TrimSpace(data)

	if bytes.Equal(trimmed, []byte("null")) {
		*l = nil
		return nil
	}

	// Plain array of strings.
	var arr []string
	if err := json.Unmarshal(trimmed, &arr); err == nil {
		*l = arr
		return nil
	}

	// Raw CSV string (defensive: some endpoints return the joined list).
	var str string
	if err := json.Unmarshal(trimmed, &str); err == nil {
		*l = splitCSVList(str)
		return nil
	}

	// Object form: {"00..9": value} (index -> value) or {"key": value}
	// (MultiSelect value -> value / relation UUID -> label).
	var objStr map[string]string
	if err := json.Unmarshal(trimmed, &objStr); err == nil {
		*l = stringListFromMap(objStr)
		return nil
	}

	// Object with non-string values: treat the keys as the values.
	var objAny map[string]json.RawMessage
	if err := json.Unmarshal(trimmed, &objAny); err == nil {
		keys := make([]string, 0, len(objAny))
		for k := range objAny {
			keys = append(keys, k)
		}
		// keys are unordered; reuse the map helper for a deterministic result
		m := make(map[string]string, len(keys))
		for _, k := range keys {
			m[k] = k
		}
		*l = stringListFromMap(m)
		return nil
	}

	return fmt.Errorf("StringList: unsupported JSON value %s", trimmed)
}

// stringListFromMap converts an object-shaped list of strings to a slice.
// Keys that are all pure integers are treated as indexes into the values
// ({"0":"a","1":"b"} -> ["a","b"], ordered by index); otherwise the keys
// themselves are the values (MultiSelect value->value, relation UUID->label),
// sorted for a deterministic result.
func stringListFromMap(m map[string]string) StringList {
	if len(m) == 0 {
		return StringList{}
	}

	keys := make([]string, 0, len(m))
	allInteger := true
	for k := range m {
		keys = append(keys, k)
		if _, err := strconv.Atoi(k); err != nil {
			allInteger = false
		}
	}

	out := make(StringList, 0, len(m))
	if allInteger {
		sort.Slice(keys, func(i, j int) bool {
			ki, _ := strconv.Atoi(keys[i])
			kj, _ := strconv.Atoi(keys[j])
			return ki < kj
		})
		for _, k := range keys {
			out = append(out, m[k])
		}
		return out
	}

	sort.Strings(keys)
	for _, k := range keys {
		out = append(out, k)
	}
	return out
}

// splitCSVList splits a comma-separated string into a list, dropping empty
// entries.
func splitCSVList(s string) StringList {
	out := StringList{}
	for _, part := range strings.Split(s, ",") {
		if part = strings.TrimSpace(part); part != "" {
			out = append(out, part)
		}
	}
	return out
}
