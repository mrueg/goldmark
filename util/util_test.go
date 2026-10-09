package util

import (
	"testing"
	"unicode/utf8"
)

func TestToRune(t *testing.T) {
	cases := []struct {
		src  string
		pos  int
		want rune
	}{
		{"a", 0, 'a'},
		{"日本", 0, '日'},
		{"日本", 2, '日'},
		{"日本", 5, '本'},
		{"a日", 3, '日'},
		// no rune start before pos
		{"\x80", 0, utf8.RuneError},
		{"\xbf\xbf", 1, utf8.RuneError},
		// the rune found before pos does not cover pos
		{"a\x80", 1, utf8.RuneError},
		// truncated multi-byte sequence
		{"\xe6\x97", 1, utf8.RuneError},
	}
	for _, c := range cases {
		if got := ToRune([]byte(c.src), c.pos); got != c.want {
			t.Errorf("ToRune(%q, %d) = %q, want %q", c.src, c.pos, got, c.want)
		}
	}
}
