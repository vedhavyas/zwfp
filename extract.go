package zwfp

import (
	"fmt"
	"strconv"
	"strings"
)

// separate separates plaintext and zero-width embedded in the string s
func separate(s string) ([]rune, []rune) {
	var pt, zw []rune
	for _, c := range s {
		switch c {
		case zwnb, zwsp, zwj, zwnj:
			zw = append(zw, c)
		default:
			pt = append(pt, c)
		}
	}

	return pt, zw
}

// constructLetter converts zero-width to plain text letter
// zwsp -> 1
// zwnj -> 0
func constructLetter(zws []rune) (rune, error) {
	var sb strings.Builder
	for _, r := range zws {
		switch r {
		case zwsp:
			sb.WriteString("1")
		case zwnj:
			sb.WriteString("0")
		default:
			continue
		}
	}

	n, err := strconv.ParseInt(sb.String(), 2, 32)
	if err != nil {
		return 0, fmt.Errorf("failed to convert %v to letter: %v", zws, err)
	}

	return rune(n), nil
}

// constructKey constructs the embed key from zero-width string
func constructKey(zws []rune) string {
	var sb strings.Builder
	var cl []rune
	for _, r := range zws {
		switch r {
		case zwj, zwnb:
			dr, err := constructLetter(cl)
			if err != nil {
				continue
			}

			sb.WriteRune(dr)
			cl = nil
			if r == zwnb {
				sb.WriteString(" ")
			}

		default:
			cl = append(cl, r)
		}

	}

	if len(cl) > 0 {
		r, err := constructLetter(cl)
		if err == nil {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}

func Extract(embed string) (plainText, key string) {
	pr, zws := separate(embed)
	key = constructKey(zws)
	return string(pr), key
}
