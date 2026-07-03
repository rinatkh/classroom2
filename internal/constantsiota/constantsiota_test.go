package constantsiota

import "testing"

func TestExample(t *testing.T) {
	want := "app: go-course\n" +
		"max attempts: 3\n" +
		"statuses: 0 1 2\n" +
		"is paid: true\n" +
		"roles: 0 1 2\n"

	if got := Example(); got != want {
		t.Fatalf("Example() = %q, want %q", got, want)
	}
}
