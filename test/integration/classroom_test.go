package integration

import (
	"fmt"
	"os/exec"
	"testing"
	"unsafe"
)

func TestCommandOutputs(t *testing.T) {
	var i int = 10
	var i64 int64 = 10
	var flag bool = true
	var letter rune = 'Ж'

	cases := []struct {
		name string
		path string
		want string
	}{
		{
			name: "declaration",
			path: "../../cmd/01_declaration",
			want: "zero int: 0\n" +
				"lesson: 2\n" +
				"language: Go\n" +
				"student: Maria\n" +
				"after +=: 15\n" +
				"after ++: 16\n" +
				"rectangle: 10 20\n" +
				"new rectangle: 15 25\n" +
				"box: 30 25 5\n" +
				"profile: Maria 25 Berlin\n",
		},
		{
			name: "numbers",
			path: "../../cmd/02_numbers",
			want: "users: 125\n" +
				"small/big: 120 9000000000\n" +
				"attempts: 3\n" +
				"a+b: 22\n" +
				"a-b: 12\n" +
				"a*b: 85\n" +
				"a/b: 3\n" +
				"a mod b: 2\n" +
				"final price: 1050\n",
		},
		{
			name: "bases-bytes",
			path: "../../cmd/03_bases_bytes",
			want: fmt.Sprintf("same value: 42 42 42 42\ndec=255 bin=11111111 oct=377 hex=ff\nsizeof int: %d\nsizeof int64: %d\nsizeof bool: %d\nsizeof rune: %d\ndec=64 bin=1000000 hex=40\n",
				unsafe.Sizeof(i),
				unsafe.Sizeof(i64),
				unsafe.Sizeof(flag),
				unsafe.Sizeof(letter),
			),
		},
		{
			name: "float-bool-complex",
			path: "../../cmd/04_float_bool_complex",
			want: "final=239.88\n" +
				"can enter: true\n" +
				"can rest: true\n" +
				"workday: false\n" +
				"complex: (2+3i)\n" +
				"average=4.33\n" +
				"passed: true\n",
		},
		{
			name: "strings-runes",
			path: "../../cmd/05_strings_runes",
			want: "bytes: 16\n" +
				"runes: 10\n" +
				"first rune: П\n" +
				"first byte: 71\n" +
				"first byte as char: G\n" +
				"normalized: Maria\n" +
				"upper: GOLANG\n",
		},
		{
			name: "constants-iota",
			path: "../../cmd/06_constants_iota",
			want: "app: go-course\n" +
				"max attempts: 3\n" +
				"statuses: 0 1 2\n" +
				"is paid: true\n" +
				"roles: 0 1 2\n",
		},
		{
			name: "conversion-named-types",
			path: "../../cmd/07_conversion_named_types",
			want: "sum: 30\n" +
				"user id: 1001\n" +
				"kop: 120000\n" +
				"atoi: 42 <nil>\n" +
				"itoa: 42\n" +
				"result: 130\n" +
				"parse error: <nil>\n",
		},
		{
			name: "pointers",
			path: "../../cmd/08_pointers",
			want: "before: 10\n" +
				"after pointer change: 11\n" +
				"created with new: 42\n" +
				"nil pointer: true\n" +
				"default value: 100\n",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cmd := exec.Command("go", "run", tc.path)
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("go run %s failed: %v\n%s", tc.path, err, out)
			}

			if got := string(out); got != tc.want {
				t.Fatalf("output = %q, want %q", got, tc.want)
			}
		})
	}
}
