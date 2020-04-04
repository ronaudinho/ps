package util

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	rti = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	ErrInvalidRoman            = func(r string) error { return errors.New(fmt.Sprintf("invalid roman numerals %s", r)) }
	ErrUndeclaredIntergalactic = func(r string) error { return errors.New(fmt.Sprintf("undeclared intergalactic units in %s", r)) }
)

func IntergalacticToInt(intergalactic, sep string, unit map[string]string) (int, error) {
	r, err := IntergalacticToRoman(intergalactic, sep, unit)
	if err != nil {
		return 0, err
	}
	a, err := RomanToInt(r)
	if err != nil {
		return 0, err
	}
	return a, nil
}

func IntergalacticToRoman(intergalactic, sep string, unit map[string]string) (string, error) {
	var r string
	c := strings.Split(intergalactic, sep)
	for _, x := range c {
		_, ok := unit[x]
		if !ok {
			return "", ErrUndeclaredIntergalactic(intergalactic)
		}
		r = r + unit[x]
	}
	return r, nil
}

func RomanToInt(roman string) (int, error) {
	var a int
	err := ValidateRoman(roman)
	if err != nil {
		return 0, ErrInvalidRoman(roman)
	}
	for i := 0; i < len(roman); i++ {
		if i == len(roman)-1 {
			a += rti[roman[len(roman)-1]]
			break
		}
		if rti[roman[i]] >= rti[roman[i+1]] {
			a += rti[roman[i]]
			continue
		}
		a += rti[roman[i+1]] - rti[roman[i]]
		i++
	}
	return a, nil
}

func ValidateRoman(roman string) error {
	valid, _ := regexp.Match("^M{0,3}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$", []byte(roman))
	if !valid {
		return ErrInvalidRoman(roman)
	}
	return nil
}
