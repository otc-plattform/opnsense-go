package api

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestStringListUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    StringList
		wantErr bool
	}{
		{"empty object", `{}`, StringList{}, false},
		{"empty object with whitespace", `{  }`, StringList{}, false},
		{"empty array", `[]`, StringList{}, false},
		{"null", `null`, nil, false},
		{"array", `["127.0.0.1","::1"]`, StringList{"127.0.0.1", "::1"}, false},
		{"empty placeholder", `{"":{"value":"","selected":1}}`, StringList{}, false},
		{"options", `{"10.0.0.2":{"value":"10.0.0.2","selected":true},"10.0.0.1":{"value":"10.0.0.1","selected":1},"10.0.0.3":{"selected":0}}`, StringList{"10.0.0.1", "10.0.0.2"}, false},
		{"invalid object", `{"address":"127.0.0.1"}`, nil, true},
		{"string", `"127.0.0.1"`, nil, true},
		{"invalid array", `[1]`, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StringList{"old"}
			err := json.Unmarshal([]byte(tt.input), &got)
			if (err != nil) != tt.wantErr {
				t.Fatalf("error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %#v, want %#v", got, tt.want)
			}
		})
	}
}
