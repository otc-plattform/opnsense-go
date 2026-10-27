package api

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestStringListUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		want    StringList
		wantErr bool
	}{
		{name: "null", json: `null`, want: nil},
		{name: "empty array", json: `[]`, want: StringList{}},
		{name: "array", json: `["a","b"]`, want: StringList{"a", "b"}},
		{name: "map key equals value", json: `{"127.0.0.1":"127.0.0.1"}`, want: StringList{"127.0.0.1"}},
		{name: "map index to value", json: `{"1":"b","0":"a"}`, want: StringList{"a", "b"}},
		{name: "map uuid key to label (relation)", json: `{"6cf7..uuid":"http-80"}`, want: StringList{"6cf7..uuid"}},
		{name: "csv string", json: `"a,b"`, want: StringList{"a", "b"}},
		{name: "empty csv string", json: `""`, want: StringList{}},
		{name: "single string", json: `"http-80"`, want: StringList{"http-80"}},
		{name: "map with bool values", json: `{"http-80":true}`, want: StringList{"http-80"}},
		{name: "empty object", json: `{}`, want: StringList{}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var got StringList
			err := json.Unmarshal([]byte(tc.json), &got)
			if tc.wantErr != (err != nil) {
				t.Fatalf("error mismatch: got err=%v, wantErr=%v", err, tc.wantErr)
			}
			if err == nil && !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("value mismatch: got %v, want %v", got, tc.want)
			}
		})
	}
}
