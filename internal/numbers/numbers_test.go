package numbers

import "testing"

func TestExample(t *testing.T) {
	want := "users: 125\n" +
		"small/big: 120 9000000000\n" +
		"attempts: 3\n" +
		"a+b: 22\n" +
		"a-b: 12\n" +
		"a*b: 85\n" +
		"a/b: 3\n" +
		"a mod b: 2\n" +
		"final price: 1050\n"

	if got := Example(); got != want {
		t.Fatalf("Example() = %q, want %q", got, want)
	}
}
