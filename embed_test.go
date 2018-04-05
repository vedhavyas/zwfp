package zwfp

import (
	"reflect"
	"testing"
)

func Test_toBits(t *testing.T) {
	tests := []struct {
		s string
		r []string
	}{
		{
			s: "s",
			r: []string{"1110011"},
		},

		{
			s: "Hello\u200D",
			r: []string{
				"1001000",
				"1100101",
				"1101100",
				"1101100",
				"1101111",
				"10000000001101",
			},
		},
	}

	for _, c := range tests {
		r := toBits(c.s)
		if !reflect.DeepEqual(c.r, r) {
			t.Fatalf("expected %v bits bit got %v", c.r, r)
		}
	}
}

func test(t *testing.T, tests []struct{ s, r string }, fn func(s string) string) {
	for _, c := range tests {
		r := fn(c.s)
		if r != c.r {
			t.Fatalf("expected %s but got %s", c.r, r)
		}
	}
}

func Test_convertLetter(t *testing.T) {
	tests := []struct {
		s string
		r string
	}{
		{
			s: "1101100",
			r: "\u200B\u200B\u200C\u200B\u200B\u200C\u200C",
		},

		{
			s: "1101111",
			r: "\u200B\u200B\u200C\u200B\u200B\u200B\u200B",
		},
	}

	test(t, tests, convertLetter)
}

func Test_convertWord(t *testing.T) {
	tests := []struct {
		s string
		r string
	}{
		{
			s: "h",
			r: "\u200B\u200B\u200C\u200B\u200C\u200C\u200C",
		},

		{
			s: "hi",
			r: "\u200B\u200B\u200C\u200B\u200C\u200C\u200C\u200D\u200B\u200B\u200C\u200B\u200C\u200C\u200B",
		},
	}

	test(t, tests, convertWord)
}

func Test_toZeroWidth(t *testing.T) {
	tests := []struct {
		s string
		r string
	}{
		{
			s: "hi",
			r: "\u200B\u200B\u200C\u200B\u200C\u200C\u200C\u200D\u200B\u200B\u200C\u200B\u200C\u200C\u200B",
		},

		{
			s: "hi ",
			r: "\u200B\u200B\u200C\u200B\u200C\u200C\u200C\u200D\u200B\u200B\u200C\u200B\u200C\u200C\u200B",
		},

		{
			s: "hi wd",
			r: "\u200B\u200B\u200C\u200B\u200C\u200C\u200C\u200D\u200B\u200B\u200C\u200B\u200C\u200C\u200B\uFEFF\u200B\u200B\u200B\u200C\u200B\u200B\u200B\u200D\u200B\u200B\u200C\u200C\u200B\u200C\u200C",
		},
	}

	test(t, tests, toZeroWidth)
}

func TestEmbed(t *testing.T) {
	tests := []struct {
		s string
		k string
		r string
	}{
		{
			s: "Hello",
			k: "hi",
			r: "H\u200Be\u200Bl\u200Cl\u200Bo\u200C\u200C\u200C\u200D\u200B\u200B\u200C\u200B\u200C\u200C\u200B",
		},

		{
			s: "Hello, World!",
			k: "hi",
			r: "H\u200Be\u200Bl\u200Cl\u200Bo\u200C,\u200C \u200CW\u200Do\u200Br\u200Bl\u200Cd\u200B!\u200C\u200C\u200B",
		},

		{
			s: "",
			k: "hi",
			r: "\u200B\u200B\u200C\u200B\u200C\u200C\u200C\u200D\u200B\u200B\u200C\u200B\u200C\u200C\u200B",
		},

		{
			s: "",
			k: "",
			r: "",
		},

		{
			s: "Hello, World!!!",
			k: "",
			r: "Hello, World!!!",
		},

		{
			s: "eat",
			k: "hi wd",
			r: "e\u200Ba\u200Bt\u200C\u200B\u200C\u200C\u200C\u200D\u200B\u200B\u200C\u200B\u200C\u200C\u200B\uFEFF\u200B\u200B\u200B\u200C\u200B\u200B\u200B\u200D\u200B\u200B\u200C\u200C\u200B\u200C\u200C",
		},
	}

	for _, c := range tests {
		r := Embed(c.s, c.k)
		if r != c.r {
			t.Fatalf("expected %s but got %s", c.r, r)
		}
	}
}
