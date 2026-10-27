package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
)

// StringList decodes string arrays and OPNsense option maps, including the
// empty object and empty-key placeholder sometimes returned for empty lists.
type StringList []string

func (s *StringList) UnmarshalJSON(data []byte) error {
	if bytes.Equal(bytes.TrimSpace(data), []byte("null")) {
		*s = nil
		return nil
	}

	var values []string
	if err := json.Unmarshal(data, &values); err == nil {
		*s = values
		return nil
	}

	var object map[string]struct {
		Selected json.RawMessage `json:"selected"`
	}
	if err := json.Unmarshal(data, &object); err != nil {
		return fmt.Errorf("expected a string array or option map: %w", err)
	}
	values = []string{}
	for key, option := range object {
		switch string(option.Selected) {
		case "1", "true":
			if key != "" {
				values = append(values, key)
			}
		case "0", "false":
		default:
			return fmt.Errorf("invalid selection for %q", key)
		}
	}
	sort.Strings(values)
	*s = values
	return nil
}
