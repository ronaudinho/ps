package ps

import (
	"regexp"

	"github.com/ronaudinho/ps/pkg/util"
)

type Unit struct {
	Intergalactic string `key:"intergalactic"`
	Roman         string `key:"roman"`
}

// UnitParser implements Parser interface
type UnitParser struct {
	exp   string
	regex *regexp.Regexp
}

// NewUnitParser initiates an instance of UnitParser
func NewUnitParser(exp string) *UnitParser {
	return &UnitParser{
		exp:   exp,
		regex: regexp.MustCompile(exp),
	}
}

// Parse parses line, looking for intergalactic and roman mapping
// returns an error if roman is invalid or mapping already exists
// uses struct to store regex result
func (up *UnitParser) Parse(line string, store Store) error {
	m := util.MapRegex(line, up.regex)
	var u Unit
	util.Unmarshal(m, &u)
	err := util.ValidateRoman(u.Roman)
	if err != nil {
		return err
	}
	err = store.Add("unit", u.Intergalactic, u.Roman)
	if err != nil {
		return err
	}
	return nil
}

func (up *UnitParser) regexp() string {
	return up.exp
}

func (up *UnitParser) Strout() string {
	return ""
}

// UnitMapParser implements Parser interface
type UnitMapParser struct {
	exp   string
	regex *regexp.Regexp
}

// NewUnitMapParser initiates an instance of UnitMapParser
func NewUnitMapParser(exp string) *UnitMapParser {
	return &UnitMapParser{
		exp:   exp,
		regex: regexp.MustCompile(exp),
	}
}

// Parse parses line, looking for intergalactic and roman mapping
// returns an error if roman is invalid or mapping already exists
// uses struct to store regex result
func (up *UnitMapParser) Parse(line string, store Store) error {
	m := util.MapRegex(line, up.regex)
	err := util.ValidateRoman(m["roman"])
	if err != nil {
		return err
	}
	err = store.Add("unit", m["intergalactic"], m["roman"])
	if err != nil {
		return err
	}
	return nil
}

func (up *UnitMapParser) regexp() string {
	return up.exp
}

func (up *UnitMapParser) Strout() string {
	return ""
}
