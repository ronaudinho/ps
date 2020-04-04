package util_test

import (
	"testing"

	"github.com/ronaudinho/ps/pkg/util"
)

var (
	itr = map[string]string{
		"glob": "I",
		"prok": "V",
		"pish": "X",
		"tegj": "L",
	}
)

func TestIntergalacticToInt(t *testing.T) {
	tests := []struct {
		desc  string
		in    string
		sep   string
		itr   map[string]string
		out   int
		iserr bool
	}{
		{
			desc: "valid_intergalactic",
			in:   "glob glob",
			sep:  " ",
			itr:  itr,
			out:  2,
		},
		{
			desc:  "invalid_intergalactic_undeclared_unit",
			in:    "glob blog",
			sep:   " ",
			itr:   itr,
			out:   0,
			iserr: true,
		},
		{
			desc:  "invalid_intergalactic_invalid_roman",
			in:    "prok prok",
			sep:   " ",
			itr:   itr,
			out:   0,
			iserr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			a, err := util.IntergalacticToInt(tt.in, tt.sep, tt.itr)
			if err != nil && tt.iserr != true {
				t.Errorf("want %v got %s", nil, err.Error())
			}
			if a != tt.out {
				t.Errorf("want %d got %d", tt.out, a)
			}
		})
	}
}

func TestIntergalacticToRoman(t *testing.T) {
	tests := []struct {
		desc  string
		in    string
		sep   string
		itr   map[string]string
		out   string
		iserr bool
	}{
		{
			desc: "valid_intergalactic",
			in:   "glob glob",
			sep:  " ",
			itr:  itr,
			out:  "II",
		},
		{
			desc:  "invalid_intergalactic",
			in:    "globglob",
			sep:   " ",
			itr:   itr,
			out:   "",
			iserr: true,
		},
		{
			desc:  "invalid_intergalactic",
			in:    "glob blog",
			sep:   " ",
			itr:   itr,
			out:   "",
			iserr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			r, err := util.IntergalacticToRoman(tt.in, tt.sep, tt.itr)
			if err != nil && tt.iserr != true {
				t.Errorf("want %v got %s", nil, err.Error())
			}
			if r != tt.out {
				t.Errorf("want %s got %s", tt.out, r)
			}
		})
	}
}

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		desc  string
		in    string
		out   int
		iserr bool
	}{
		{
			desc: "valid_roman",
			in:   "XXXIX",
			out:  39,
		},
		{
			desc:  "invalid_roman",
			in:    "Z",
			out:   0,
			iserr: true,
		},
		{
			desc:  "invalid_too_many_repetition_IXCM",
			in:    "XXXX",
			out:   0,
			iserr: true,
		},
		{
			desc:  "invalid_too_many_repetition_DLV",
			in:    "VV",
			out:   0,
			iserr: true,
		},
		{
			desc:  "invalid_substraction_IXC",
			in:    "IL",
			out:   0,
			iserr: true,
		},
		{
			desc:  "invalid_substraction_VLD",
			in:    "VX",
			out:   0,
			iserr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			a, err := util.RomanToInt(tt.in)
			if err != nil && tt.iserr != true {
				t.Errorf("want %v got %s", nil, err.Error())
			}
			if a != tt.out {
				t.Errorf("want %d got %d", tt.out, a)
			}
		})
	}
}
