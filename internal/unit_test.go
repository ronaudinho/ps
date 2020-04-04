package ps_test

import (
	"testing"

	"github.com/ronaudinho/ps/internal"
)

var (
	DefUnitIn = []string{
		"glob is I",
		"prok is V",
		"pish is X",
		"tegj is L",
	}
	DupRomanIn = []string{
		"glob is I",
		"prok is I",
	}
	DupIntergalacticIn = []string{
		"glob is I",
		"glob is V",
	}
	InvRomanFmt = []string{
		"glob is IL",
	}
)

func TestUnitParser_Parse(t *testing.T) {
	tests := []struct {
		desc  string
		in    []string
		iserr bool
	}{
		{
			desc:  "default",
			in:    DefUnitIn,
			iserr: false,
		},
		{
			desc:  "duplicate_roman_unit",
			in:    DupRomanIn,
			iserr: false,
		},
		{
			desc:  "duplicate_intergalactic_unit",
			in:    DupIntergalacticIn,
			iserr: true,
		},
		{
			desc:  "invalid_roman_format",
			in:    InvRomanFmt,
			iserr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			p := ps.NewUnitParser(`(?P<intergalactic>[a-z]+) is (?P<roman>[IVXLCDM]+)`)
			kv := ps.KVStore(make(map[string]map[string]interface{}))
			for _, u := range tt.in {
				err := p.Parse(u, kv)
				if err != nil && tt.iserr != true {
					t.Errorf("want %v got %s", nil, err.Error())
				}
			}
		})
	}
}

func BenchmarkUnitParser_Parse(b *testing.B) {
	tests := []struct {
		desc  string
		in    []string
		iserr bool
	}{
		{
			desc:  "default",
			in:    DefUnitIn,
			iserr: false,
		},
	}
	for _, tt := range tests {
		b.Run(tt.desc, func(b *testing.B) {
			p := ps.NewUnitParser(`(?P<intergalactic>[a-z]+) is (?P<roman>[IVXLCDM]+)`)
			kv := ps.KVStore(make(map[string]map[string]interface{}))
			for _, u := range tt.in {
				p.Parse(u, kv)
			}
		})
	}
}

func BenchmarkUnitMapParser_Parse(b *testing.B) {
	tests := []struct {
		desc  string
		in    []string
		iserr bool
	}{
		{
			desc:  "default",
			in:    DefUnitIn,
			iserr: false,
		},
	}
	for _, tt := range tests {
		b.Run(tt.desc, func(b *testing.B) {
			p := ps.NewUnitMapParser(`(?P<intergalactic>[a-z]+) is (?P<roman>[IVXLCDM]+)`)
			kv := ps.KVStore(make(map[string]map[string]interface{}))
			for _, u := range tt.in {
				p.Parse(u, kv)
			}
		})
	}
}
