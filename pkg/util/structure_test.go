package util_test

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/ronaudinho/ps/pkg/util"
)

var (
	reg = regexp.MustCompile(`^(?P<intergalactic>[a-z]*) is (?P<roman>[IVXLCDM]*)`)
)

func TestMapRegex(t *testing.T) {
	tests := []struct {
		desc string
		in   string
		out  map[string]string
	}{
		{
			desc: "default",
			in:   "glob is I",
			out: map[string]string{
				"intergalactic": "glob",
				"roman":         "I",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			m := util.MapRegex(tt.in, reg)
			if !reflect.DeepEqual(tt.out, m) {
				t.Errorf("want %v got %s", tt.out, m)
			}
		})
	}
}

func TestUnmarshal(t *testing.T) {
	tests := []struct {
		desc string
		in   map[string]string
		out  interface{}
	}{
		{
			desc: "default",
			in: map[string]string{
				"intergalactic": "glob",
				"roman":         "I",
			},
			out: struct {
				Intergalactic string `key:"intergalactic"`
				Roman         string `key:"roman"`
			}{
				Intergalactic: "glob",
				Roman:         "I",
			},
		},
		{
			desc: "wrong_key_name",
			in: map[string]string{
				"intergalactic": "glob",
				"romanista":     "I",
			},
			out: struct {
				Intergalactic string `key:"intergalactic"`
				Roman         string `key:"roman"`
			}{
				Intergalactic: "glob",
			},
		},
		{
			desc: "non_existent_key",
			in: map[string]string{
				"intergalactic": "glob",
				"roman":         "I",
				"arabic":        "1",
			},
			out: struct {
				Intergalactic string `key:"intergalactic"`
				Roman         string `key:"roman"`
			}{
				Intergalactic: "glob",
				Roman:         "I",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			o := struct {
				Intergalactic string `key:"intergalactic"`
				Roman         string `key:"roman"`
			}{}
			util.Unmarshal(tt.in, &o)
			if !reflect.DeepEqual(tt.out, o) {
				t.Errorf("want %v got %s", tt.out, o)
			}
		})
	}
}
