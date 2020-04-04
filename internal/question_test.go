package ps_test

import (
	"testing"

	"github.com/ronaudinho/ps/internal"
)

const (
	DefAns = "IDK LOL"
)

// should probably define store mock somewhere else
var (
	DefKVPrice = map[string]interface{}{
		"Silver": float32(17),
		"Gold":   float32(14450),
		"Iron":   float32(195.5),
	}
)

func TestQuestionParser_Parse(t *testing.T) {
	tests := []struct {
		desc string
		in   []string
		out  string
	}{
		{
			desc: "invalid_question",
			in: []string{
				"how do you love me ?",
			},
			out: DefAns,
		},
		{
			desc: "how_much",
			in: []string{
				"how much is pish tegj glob glob ?",
			},
			out: "pish tegj glob glob is 42",
		},
		{
			desc: "how_much_undefined_intergalactic",
			in: []string{
				"how much is pish tegj glob blog ?",
			},
			out: DefAns,
		},
		{
			desc: "how_much_invalid_intergalactic",
			in: []string{
				"how much is glob tegj ?",
			},
			out: DefAns,
		},
		{
			desc: "how_many",
			in: []string{
				"how many Credits is glob prok Silver ?",
			},
			out: "glob prok Silver is 68 Credits",
		},
		{
			desc: "how_many_undefined_intergalactic",
			in: []string{
				"how many Credits is blog prok Silver ?",
			},
			out: DefAns,
		},
		{
			desc: "how_many_invalid_intergalactic",
			in: []string{
				"how many Credits is glob tegj  Silver ?",
			},
			out: DefAns,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			for _, l := range tt.in {
				p := ps.NewQuestionParser(
					DefAns,
					`^([\S\s]*) \?`,
					map[string]ps.Parser{
						"howmuch": ps.NewHowMuchQuestionParser(`^how much is (?P<intergalactic>[a-z ]+) \?`),
						"howmany": ps.NewHowManyQuestionParser(`^how many Credits is (?P<intergalactic>[a-z ]+) (?P<item>[A-Z][a-z]+) \?`),
					},
				)
				kv := ps.KVStore(make(map[string]map[string]interface{}))
				kv["unit"] = DefKVUnit
				kv["price"] = DefKVPrice
				p.Parse(l, kv)
				if p.Strout() != tt.out {
					t.Errorf("want %v got %s", tt.out, p.Strout())
				}
			}
		})
	}
}
