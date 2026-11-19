package conversionnamedtypes

import "testing"

func TestExample(t *testing.T) {
	want := "sum: 30\n" +
		"user id: 1001\n" +
		"kop: 120000\n" +
		"atoi: 42 <nil>\n" +
		"itoa: 42\n" +
		"result: 130\n" +
		"parse error: <nil>\ninvalid input: abc\n"

	if got := Example(); got != want {
		t.Fatalf("Example() = %q, want %q", got, want)
	}
}
