package ps_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/ronaudinho/ps/internal"
)

const (
	InputTestsDir  = "../tests/in"
	OutputTestsDir = "../tests/out"
)

// this is probably integration test
func TestParseFile(t *testing.T) {
	tests := []struct {
		desc  string
		in    string
		out   string
		iserr bool
	}{
		{
			desc:  "default",
			in:    fmt.Sprintf("%s/default.txt", InputTestsDir),
			out:   fmt.Sprintf("%s/default.txt", OutputTestsDir),
			iserr: false,
		},
		{
			desc:  "duplicate_roman_unit",
			in:    fmt.Sprintf("%s/duplicate_roman_unit.txt", InputTestsDir),
			out:   fmt.Sprintf("%s/duplicate_roman_unit.txt", OutputTestsDir),
			iserr: false,
		},
		{
			desc:  "duplicate_intergalactic_unit",
			in:    fmt.Sprintf("%s/duplicate_intergalactic_unit.txt", InputTestsDir),
			out:   fmt.Sprintf("%s/duplicate_intergalactic_unit.txt", OutputTestsDir),
			iserr: true,
		},
		{
			desc:  "invalid_intergalactic_unit",
			in:    fmt.Sprintf("%s/invalid_intergalactic_unit.txt", InputTestsDir),
			out:   fmt.Sprintf("%s/invalid_intergalactic_unit.txt", OutputTestsDir),
			iserr: true,
		},
		{
			desc:  "invalid_roman_format",
			in:    fmt.Sprintf("%s/invalid_roman_format.txt", InputTestsDir),
			out:   fmt.Sprintf("%s/invalid_roman_format.txt", OutputTestsDir),
			iserr: true,
		},
		{
			desc:  "invalid_roman_unit",
			in:    fmt.Sprintf("%s/invalid_roman_unit.txt", InputTestsDir),
			out:   fmt.Sprintf("%s/invalid_roman_unit.txt", OutputTestsDir),
			iserr: false, // simply skipped by regex, may need another handling
		},
		{
			desc:  "undefined_item_price",
			in:    fmt.Sprintf("%s/undefined_item_price.txt", InputTestsDir),
			out:   fmt.Sprintf("%s/undefined_item_price.txt", OutputTestsDir),
			iserr: false, // if item price is undefined, fallback to default answer
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			p := ps.NewFileParser(
				map[string]ps.Parser{
					"unit":      ps.NewUnitParser(`(?P<intergalactic>[a-z]+) is (?P<roman>[IVXLCDM]+)`),
					"itemprice": ps.NewItemPriceParser(`^(?P<intergalactic>[a-z ]+) (?P<item>[A-Z][a-z]+) is (?P<price>[0-9]+) Credits`),
					"question": ps.NewQuestionParser(
						"I have no idea what you are talking about",
						`^([\S\s]*) \?`,
						map[string]ps.Parser{
							"howmuch": ps.NewHowMuchQuestionParser(`^how much is (?P<intergalactic>[a-z ]+) \?`),
							"howmany": ps.NewHowManyQuestionParser(`^how many Credits is (?P<intergalactic>[a-z ]+) (?P<item>[A-Z][a-z]+) \?`),
						},
					),
				},
				ps.KVStore(make(map[string]map[string]interface{})),
			)
			err := p.Parse(tt.in, nil)
			b, _ := ioutil.ReadFile(tt.out)
			if err == nil {
				if tt.iserr == true {
					t.Errorf("want %v got %s", nil, err.Error())
				} else if p.Strout() != string(b[:len(b)-1]) {
					t.Errorf("want:\n%s\ngot:\n%s", string(b), p.Strout())
				}
			}
			if err != nil {
				if tt.iserr != true {
					t.Errorf("want %v got %s", nil, err.Error())
				} else if tt.iserr == true && fmt.Sprintf("ERROR: %s", err.Error()) != string(b[:len(b)-1]) {
					t.Errorf("want %v got %s", fmt.Sprintf("ERROR: %s", err.Error()), string(b[:len(b)-1]))
				}
			}
		})
	}
}
