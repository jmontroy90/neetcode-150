package arrays_hashing

import (
	"fmt"
	"math"
	"strings"
)

const (
	delim = '-'
)

func Encode(strs []string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(fmt.Sprintf("%d-%v", len(str), str))
	}
	return sb.String()
}

func Decode(str string) []string {
	var out []string
	var word string
	remaining := str
	for {
		if remaining == "" {
			break
		}
		word, remaining = decodeWord(remaining)
		out = append(out, word)
	}
	return out
}

func decodeWord(str string) (string, string) {
	strRunes := []rune(str)
	var ns []rune
	var i int
	for strRunes[i] != delim {
		ns = append(ns, strRunes[i])
		i++
	}
	i++ // move past delim
	// this whole block actually parses the string numbers into an int, digit by digit
	var n int
	for ix, rn := range ns {
		ni := int(rn - '0')
		np := int(math.Pow(10.0, float64(len(ns)-1-ix)))
		n += ni * np
	}
	return string(strRunes[i : i+n]), string(strRunes[i+n:])
}
