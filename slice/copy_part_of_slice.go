package slice

import (
	"regexp"
)

var digitRegexp = regexp.MustCompile("[0-9]+")

func CopyDigits(text []byte) []byte {
	b := text
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

func AppendDigits(text []byte) []byte {
	b := text
	b = digitRegexp.Find(b)
	var c []byte
	c = append(c, b...)
	return c
}
