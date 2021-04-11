package utils

import (
	"fmt"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}

func RemoveAccents(s string) (clean string) {
	b := make([]byte, len(s))

	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	nDst, _, e := t.Transform(b, []byte(s), true)
	if e != nil {
		fmt.Println(e)
	}

	return string(b[:nDst])
}
