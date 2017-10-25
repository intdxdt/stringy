package stringy

import (
	"fmt"
	"strings"
	"log"
	"strconv"
)

//make strings unique
func Unique(vals []string) []string {
	uniqs := make([]string, 0)
	dict := make(map[string]struct{})
	sig := struct{}{}
	for _, v := range vals {
		dict[v] = sig
	}
	for k := range dict {
		uniqs = append(uniqs, k)
	}
	return uniqs
}

//Center string
func Center(s, pad string, width int) string {
	padsize := width - len(s)
	if padsize < 0 {
		return s
	}
	padsize = padsize / 2
	padding := strings.Repeat(pad, padsize)
	return padding + s + padding
}

//Serialize string
func Serialize(s string) string {
	tokens := []byte(s)
	buf := make([]string, len(tokens))
	for i, v := range tokens {
		buf[i] = fmt.Sprintf("%v", v)
	}
	return strings.Join(buf, ",")
}

//Deserialize
func Deserialize(s string) []byte {
	tokens := strings.Split(s, ",")
	buf := make([]byte, len(tokens))
	for i, v := range tokens {
		b, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln(err)
		}
		buf[i] = uint8(b)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return buf
}
