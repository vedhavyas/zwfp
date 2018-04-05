package zwfp

import "testing"

func Test_separate(t *testing.T) {
	tests := []struct {
		s  string
		pt string
		zw string
	}{
		{
			s:  "H\u200Be\u200Bl\u200Cl\u200Bo\u200C\u200C\u200C\u200D\u200B\u200B\u200C\u200B\u200C\u200C\u200B",
			pt: "Hello",
			zw: "\u200B\u200B\u200C\u200B\u200C\u200C\u200C\u200D\u200B\u200B\u200C\u200B\u200C\u200C\u200B",
		},

		{
			s:  "H\u200Be\u200Bl\u200Cl\u200Bo\u200C,\u200C \u200CW\u200Do\u200Br\u200Bl\u200Cd\u200B!\u200C\u200C\u200B",
			pt: "Hello, World!",
			zw: "\u200B\u200B\u200C\u200B\u200C\u200C\u200C\u200D\u200B\u200B\u200C\u200B\u200C\u200C\u200B",
		},

		{
			s:  "\u200B\u200B\u200C\u200B\u200C\u200C\u200C\u200D\u200B\u200B\u200C\u200B\u200C\u200C\u200B",
			zw: "\u200B\u200B\u200C\u200B\u200C\u200C\u200C\u200D\u200B\u200B\u200C\u200B\u200C\u200C\u200B",
		},

		{
			s:  "e\u200Ba\u200Bt\u200C\u200B\u200C\u200C\u200C\u200D\u200B\u200B\u200C\u200B\u200C\u200C\u200B\uFEFF\u200B\u200B\u200B\u200C\u200B\u200B\u200B\u200D\u200B\u200B\u200C\u200C\u200B\u200C\u200C",
			pt: "eat",
			zw: "\u200B\u200B\u200C\u200B\u200C\u200C\u200C\u200D\u200B\u200B\u200C\u200B\u200C\u200C\u200B\uFEFF\u200B\u200B\u200B\u200C\u200B\u200B\u200B\u200D\u200B\u200B\u200C\u200C\u200B\u200C\u200C",
		},

		{
			s:  "Hello, World!!!",
			pt: "Hello, World!!!",
		},
	}

	for _, c := range tests {
		pt, zw := separate(c.s)
		if string(pt) != c.pt {
			t.Fatalf("expected %s plain text but got %s", c.pt, string(pt))
		}

		if string(zw) != c.zw {
			t.Fatalf("expected %s zero-width but got %s", c.zw, string(zw))
		}
	}
}

func Test_deConvertLetter(t *testing.T) {
	tests := []struct {
		s   string
		r   rune
		err bool
	}{
		{
			s: "\u200B\u200B\u200C\u200B\u200C\u200C\u200C",
			r: 'h',
		},

		{
			err: true,
		},
	}

	for _, c := range tests {
		r, err := constructLetter([]rune(c.s))
		if err != nil {
			if c.err {
				continue
			}

			t.Fatalf("unexpected error: %v", err)
		}

		if r != c.r {
			t.Fatalf("expected %v but got %v", c.r, r)
		}
	}
}

func Test_constructKey(t *testing.T) {
	tests := []struct {
		s string
		r string
	}{
		{},

		{
			s: "\u200B\u200B\u200C\u200B\u200C\u200C\u200C",
			r: "h",
		},

		{
			s: "\u200B\u200B\u200C\u200B\u200C\u200C\u200C\u200D\u200B\u200B\u200C\u200B\u200C\u200C\u200B",
			r: "hi",
		},

		{
			s: "\u200B\u200B\u200B\u200C\u200B\u200B\u200B\u200D\u200B\u200B\u200C\u200C\u200B\u200C\u200C",
			r: "wd",
		},

		{
			s: "\u200B\u200B\u200C\u200B\u200C\u200C\u200C\u200D\u200B\u200B\u200C\u200B\u200C\u200C\u200B\uFEFF\u200B\u200B\u200B\u200C\u200B\u200B\u200B\u200D\u200B\u200B\u200C\u200C\u200B\u200C\u200C",
			r: "hi wd",
		},
	}

	for _, c := range tests {
		r := constructKey([]rune(c.s))
		if r != c.r {
			t.Fatalf("expected %s but got %s", c.r, r)
		}
	}
}

func TestExtract(t *testing.T) {
	tests := []struct {
		s   string
		pt  string
		key string
	}{
		{
			s:   "H\u200Be\u200Bl\u200Cl\u200Bo\u200C\u200C\u200C\u200D\u200B\u200B\u200C\u200B\u200C\u200C\u200B",
			pt:  "Hello",
			key: "hi",
		},

		{
			s:   "H\u200Be\u200Bl\u200Cl\u200Bo\u200C,\u200C \u200CW\u200Do\u200Br\u200Bl\u200Cd\u200B!\u200C\u200C\u200B",
			pt:  "Hello, World!",
			key: "hi",
		},

		{
			s:   "\u200B\u200B\u200C\u200B\u200C\u200C\u200C\u200D\u200B\u200B\u200C\u200B\u200C\u200C\u200B",
			pt:  "",
			key: "hi",
		},

		{
			s:   "",
			pt:  "",
			key: "",
		},

		{
			s:   "Hello, World!!!",
			pt:  "Hello, World!!!",
			key: "",
		},

		{
			s:   "e\u200Ba\u200Bt\u200C\u200B\u200C\u200C\u200C\u200D\u200B\u200B\u200C\u200B\u200C\u200C\u200B\uFEFF\u200B\u200B\u200B\u200C\u200B\u200B\u200B\u200D\u200B\u200B\u200C\u200C\u200B\u200C\u200C",
			pt:  "eat",
			key: "hi wd",
		},
	}

	for _, c := range tests {
		pt, key := Extract(c.s)
		if pt != c.pt {
			t.Fatalf("expected plain text %s but got %s", c.pt, pt)
		}

		if key != c.key {
			t.Fatalf("expected key %s but got %s", c.key, key)
		}
	}
}
