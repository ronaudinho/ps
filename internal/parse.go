package ps

import (
	"bufio"
	"io"
	"os"
	"regexp"
)

// Parser is an interface for regex parser
type Parser interface {
	Parse(string, Store) error // parse string input ([]byte would probably be more generic), saving result in Store
	regexp() string
	Strout() string
}

// FileParser implements Parser interface
type FileParser struct {
	tpl    map[string]Parser
	strict bool
	store  Store
	out    string
}

// NewFileParser initiates an instance of FileParser
func NewFileParser(tpl map[string]Parser, store Store) *FileParser {
	return &FileParser{
		tpl:   tpl,
		store: store,
	}
}

// Parse parses input file
func (fp *FileParser) Parse(in string, sto Store) error {
	f, err := os.Open(in)
	if err != nil {
		return err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		err = fp.parseLine(s)
		if err != nil {
			return err
		}
	}
	if err != nil && err != io.EOF {
		return err
	}
	return nil
}

// parseLine is a helper for parsing a line of string
func (fp *FileParser) parseLine(line string) error {
	for k, v := range fp.tpl {
		reg, err := regexp.Compile(v.regexp())
		if err != nil {
			return err
		}
		if reg.MatchString(line) {
			err = fp.tpl[k].Parse(line, fp.store)
			if err != nil {
				return err
			}
			if fp.tpl[k].Strout() != "" {
				fp.out = fp.tpl[k].Strout()
			}
		}
	}
	return nil
}

func (fp *FileParser) regexp() string {
	return ""
}

func (fp *FileParser) Strout() string {
	return fp.out
}
