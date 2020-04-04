package ps

import (
	"regexp"
	"strconv"

	"github.com/ronaudinho/ps/pkg/util"
)

type ItemPrice struct {
	Intergalactic string `key:"intergalactic"`
	Item          string `key:"item"`
	Price         int    `key:"price"`
}

// ItemPriceParser implements Parser interface
type ItemPriceParser struct {
	exp   string
	regex *regexp.Regexp
}

// NewItemPriceParser initiates an instance of ItemPriceParser
func NewItemPriceParser(exp string) *ItemPriceParser {
	return &ItemPriceParser{
		exp:   exp,
		regex: regexp.MustCompile(exp),
	}
}

// Parse parses line, looking for price of a given item
// returns an error if intergalactic to arabic conversion fails or item already exists
// uses struct to store regex result
func (ipp *ItemPriceParser) Parse(line string, store Store) error {
	m := util.MapRegex(line, ipp.regex)
	var ip ItemPrice
	util.Unmarshal(m, &ip)
	u, _ := store.Get([]string{"unit"}...)
	uu, _ := u.(map[string]interface{})
	um := make(map[string]string, len(uu))
	for k, v := range uu {
		vv, _ := v.(string)
		um[k] = vv
	}
	a, err := util.IntergalacticToInt(ip.Intergalactic, " ", um)
	if err != nil {
		return err
	}
	err = store.Add("price", ip.Item, float32(ip.Price)/float32(a))
	if err != nil {
		return err
	}
	return nil
}

func (ipp *ItemPriceParser) regexp() string {
	return ipp.exp
}

func (ipp *ItemPriceParser) Strout() string {
	return ""
}

// ItemPriceMapParser implements Parser interface
type ItemPriceMapParser struct {
	exp   string
	regex *regexp.Regexp
}

// NewItemPriceMapParser initiates an instance of ItemPriceMapParser
func NewItemPriceMapParser(exp string) *ItemPriceMapParser {
	return &ItemPriceMapParser{
		exp:   exp,
		regex: regexp.MustCompile(exp),
	}
}

// Parse parses line, looking for intergalactic value to convert to arabic numeral
// returns an error if conversion fails or item already exists
// uses map to store regex result
func (ipp *ItemPriceMapParser) Parse(line string, store Store) error {
	m := util.MapRegex(line, ipp.regex)
	u, _ := store.Get([]string{"unit"}...)
	uu, _ := u.(map[string]interface{})
	um := make(map[string]string, len(uu))
	for k, v := range uu {
		vv, _ := v.(string)
		um[k] = vv
	}
	a, _ := util.IntergalacticToInt(m["intergalactic"], " ", um)
	p, _ := strconv.ParseFloat(m["price"], 32)
	err := store.Add("price", m["item"], float32(p)/float32(a))
	if err != nil {
		return err
	}
	return nil
}

func (ipp *ItemPriceMapParser) regexp() string {
	return ipp.exp
}

func (ipp *ItemPriceMapParser) Strout() string {
	return ""
}
