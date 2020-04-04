package ps_test

import (
	"reflect"
	"testing"

	"github.com/ronaudinho/ps/internal"
)

var (
	kvStore = ps.KVStore(make(map[string]map[string]interface{}))
)

func TestStore_Add(t *testing.T) {
	tests := []struct {
		desc  string
		t     string
		k     string
		v     interface{}
		iserr bool
	}{
		{
			desc:  "table_and_first_key",
			t:     "unit",
			k:     "glob",
			v:     "I",
			iserr: false,
		},
		{
			desc:  "duplicate_key",
			t:     "unit",
			k:     "glob",
			v:     "V",
			iserr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			err := kvStore.Add(tt.t, tt.k, tt.v)
			if err != nil && tt.iserr != true {
				t.Errorf("want %v got %s", nil, err.Error())
			}
		})
	}
}

func TestStore_Get(t *testing.T) {
	tests := []struct {
		desc  string
		in    []string
		out   interface{}
		iserr bool
	}{
		{
			desc: "existing_table",
			in:   []string{"unit"},
			out: map[string]interface{}{
				"glob": "I",
			},
			iserr: false,
		},
		{
			desc:  "existing_key_existing_table",
			in:    []string{"unit", "glob"},
			out:   "I",
			iserr: false,
		},
		{
			desc:  "nonexistent_table",
			in:    []string{"price"},
			out:   nil,
			iserr: true,
		},
		{
			desc:  "nonexistent_key_existing_table",
			in:    []string{"unit", "blog"},
			out:   nil,
			iserr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			i, err := kvStore.Get(tt.in...)
			if err != nil && tt.iserr != true {
				t.Errorf("want %v got %s", nil, err.Error())
			}
			if !reflect.DeepEqual(i, tt.out) {
				t.Errorf("want %v got %v", tt.out, i)
			}
		})
	}
}
