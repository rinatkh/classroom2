package declaration

import "testing"

func TestExample(t *testing.T) {
	want := "zero int: 0\n" +
		"lesson: 2\n" +
		"language: Go\n" +
		"student: Maria\n" +
		"after +=: 15\n" +
		"after ++: 16\n" +
		"rectangle: 10 20\n" +
		"new rectangle: 15 25\n" +
		"box: 30 25 5\n" +
		"profile: Maria 25 Berlin\n"

	if got := Example(); got != want {
		t.Fatalf("Example() = %q, want %q", got, want)
	}
}
