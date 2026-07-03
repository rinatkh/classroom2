package pointers

import "testing"

func TestExample(t *testing.T) {
	want := "before: 10\n" +
		"after pointer change: 11\n" +
		"created with new: 42\n" +
		"nil pointer: true\n" +
		"default value: 100\n"

	if got := Example(); got != want {
		t.Fatalf("Example() = %q, want %q", got, want)
	}
}
