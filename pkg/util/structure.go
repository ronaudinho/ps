package util

import (
	"reflect"
	"regexp"
	"strconv"
)

// MapRegex maps subexpression names with their regex matches
func MapRegex(line string, reg *regexp.Regexp) map[string]string {
	m := make(map[string]string)
	sm := reg.FindStringSubmatch(line)
	for i, n := range reg.SubexpNames() {
		if i == 0 {
			continue
		}
		m[n] = sm[i]
	}
	return m
}

// Unmarshal unmarshals map[string]string into a struct
// matches key in map with tag 'key' in struct
// only handles string and int but can be easily extended to handle other types
func Unmarshal(in map[string]string, out interface{}) {
	t := reflect.TypeOf(out).Elem()
	for i := 0; i < t.NumField(); i++ {
		k := t.Field(i).Tag.Get("key")
		_, ok := in[k]
		if ok {
			switch t.Field(i).Type.Kind() {
			case reflect.String:
				reflect.ValueOf(out).Elem().Field(i).SetString(in[k])
			case reflect.Int:
				a, _ := strconv.Atoi(in[k])
				reflect.ValueOf(out).Elem().Field(i).SetInt(int64(a))
			}
		}
	}
}
