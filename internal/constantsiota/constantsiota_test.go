package constantsiota

import "testing"

func TestExample(t *testing.T) {
	want := "app: go-course\n" +
		"max attempts: 3\n" +
		"statuses: 0 1 2\n" +
		"is paid: true\nstatus text: paid\n" +
		"roles: 0 1 2\n"

	if got := Example(); got != want {
		t.Fatalf("Example() = %q, want %q", got, want)
	}
}

func TestStatusText(t *testing.T) {
	for _, tc := range []struct {
		name   string
		status int
		want   string
	}{
		{"new", 0, "new"}, {"paid", 1, "paid"}, {"canceled", 2, "canceled"},
		{"negative", -1, "unknown"}, {"next", 3, "unknown"}, {"large", 999, "unknown"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := StatusText(tc.status); got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}
