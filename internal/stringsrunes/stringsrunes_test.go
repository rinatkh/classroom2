package stringsrunes

import "testing"

func TestExample(t *testing.T) {
	want := "bytes: 16\n" +
		"runes: 10\n" +
		"first rune: П\n" +
		"first byte: 71\n" +
		"first byte as char: G\n" +
		"normalized: Maria\n" +
		"upper: GOLANG\n"

	if got := Example(); got != want {
		t.Fatalf("Example() = %q, want %q", got, want)
	}
}
