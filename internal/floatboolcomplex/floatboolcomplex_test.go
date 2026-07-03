package floatboolcomplex

import "testing"

func TestExample(t *testing.T) {
	want := "final=239.88\n" +
		"can enter: true\n" +
		"can rest: true\n" +
		"workday: false\n" +
		"complex: (2+3i)\n" +
		"average=4.33\n" +
		"passed: true\n"

	if got := Example(); got != want {
		t.Fatalf("Example() = %q, want %q", got, want)
	}
}
