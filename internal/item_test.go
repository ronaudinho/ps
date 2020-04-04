package ps_test

import (
	"testing"

	"github.com/ronaudinho/ps/internal"
)

var (
	// should probably define store mock somewhere else
	DefKVUnit = map[string]interface{}{
		"glob": "I",
		"prok": "V",
		"pish": "X",
		"tegj": "L",
	}
	DefItemIn = []string{
		"glob glob Silver is 34 Credits",
		"glob prok Gold is 57800 Credits",
		"pish pish Iron is 3910 Credits",
	}
)

func TestItemPriceParser_Parse(t *testing.T) {
	tests := []struct {
		desc  string
		in    []string
		iserr bool
	}{
		{
			desc:  "default",
			in:    DefItemIn,
			iserr: false,
		},
	}
	p := ps.NewItemPriceParser(`^(?P<intergalactic>[a-z ]+) (?P<item>[A-Z][a-z]+) is (?P<price>[0-9]+) Credits`)
	kv := ps.KVStore(make(map[string]map[string]interface{}))
	kv["unit"] = DefKVUnit
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			for _, u := range tt.in {
				err := p.Parse(u, kv)
				if err != nil && tt.iserr != true {
					t.Errorf("want %v got %s", nil, err.Error())
				}
			}
		})
	}
}

func BenchmarkItemPriceParser_Parse(b *testing.B) {
	tests := []struct {
		desc  string
		in    []string
		iserr bool
	}{
		{
			desc:  "default",
			in:    DefItemIn,
			iserr: false,
		},
	}
	for _, tt := range tests {
		b.Run(tt.desc, func(b *testing.B) {
			p := ps.NewItemPriceParser(`^(?P<intergalactic>[a-z ]+) (?P<item>[A-Z][a-z]+) is (?P<price>[0-9]+) Credits`)
			kv := ps.KVStore(make(map[string]map[string]interface{}))
			kv["unit"] = DefKVUnit
			for _, u := range tt.in {
				p.Parse(u, kv)
			}
		})
	}
}

func BenchmarkItemPriceMapParser_Parse(b *testing.B) {
	tests := []struct {
		desc  string
		in    []string
		iserr bool
	}{
		{
			desc:  "default",
			in:    DefItemIn,
			iserr: false,
		},
	}
	for _, tt := range tests {
		b.Run(tt.desc, func(b *testing.B) {
			p := ps.NewItemPriceMapParser(`^(?P<intergalactic>[a-z ]+) (?P<item>[A-Z][a-z]+) is (?P<price>[0-9]+) Credits`)
			kv := ps.KVStore(make(map[string]map[string]interface{}))
			kv["unit"] = DefKVUnit
			for _, u := range tt.in {
				p.Parse(u, kv)
			}
		})
	}
}
