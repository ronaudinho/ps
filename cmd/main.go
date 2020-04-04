package main

import (
	"fmt"
	"os"

	"github.com/ronaudinho/ps/internal"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please provide name of the file to open")
		os.Exit(1)
	}
	kvStore := ps.KVStore(make(map[string]map[string]interface{}))
	p := ps.NewFileParser(
		map[string]ps.Parser{
			"unit":      ps.NewUnitParser(`^(?P<intergalactic>[a-z]*) is (?P<roman>[IVXLCDM]*)`),
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
		kvStore,
	)
	err := p.Parse(os.Args[1], nil)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		// output results to stdout
		fmt.Println(p.Strout())
	}
}
