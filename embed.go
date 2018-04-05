package zwfp

import (
	"fmt"
	"strings"
)

// Zero width non-printing characters
const (
	zwsp = '\u200B' // 1
	zwnj = '\u200C' // 0
	zwj  = '\u200D' // letter
	zwnb = '\uFEFF' // word
)

// toBits converts each character in the string to base 2 form
func toBits(s string) []string {
	var bits []string
	for _, c := range s {
		bits = append(bits, fmt.Sprintf("%b", c))
	}

	return bits
}

// convert converts binary string to zero-width string
// 1 -> zwsp
// 0 -> zwnj
func convertLetter(s string) string {
	var sb strings.Builder
	for _, c := range s {
		if c == '0' {
			sb.WriteRune(zwnj)
			continue
		}

		sb.WriteRune(zwsp)
	}

	return sb.String()
}

// convertWord converts a word to zero-width string
func convertWord(s string) string {
	bits := toBits(s)
	var zws []string
	for _, b := range bits {
		zws = append(zws, convertLetter(b))
	}

	// join each letter by zero-width joiner
	return strings.Join(zws, string(zwj))
}

// toZeroWidth converts string s to zero width characters
func toZeroWidth(s string) string {
	// trim spaces across edges
	s = strings.TrimSpace(s)

	// split to words separated by space
	words := strings.Split(s, " ")

	var zwWords []string
	for _, w := range words {
		zwWords = append(zwWords, convertWord(w))
	}

	// join each word by zero-width no break
	return strings.Join(zwWords, string(zwnb))
}

// Embed embeds the key into data using zero-width characters
func Embed(data, key string) string {
	zwKey := toZeroWidth(key)
	var zwRKey []rune
	for _, c := range zwKey {
		zwRKey = append(zwRKey, c)
	}

	var t int
	var embed []rune

	for i, c := range data {
		if i == 0 {
			embed = append(embed, c)
		}

		if t < len(zwRKey) {
			embed = append(embed, zwRKey[t])
			t++
		}

		if i != 0 {
			embed = append(embed, c)
		}
	}

	if t < len(zwRKey) {
		if len(embed) > 0 {
			lb := embed[len(embed)-1]
			embed = append(embed[:len(embed)-1], zwRKey[t:]...)
			embed = append(embed, lb)
		} else {
			embed = append(embed, zwRKey[t:]...)
		}

	}

	return string(embed)
}
