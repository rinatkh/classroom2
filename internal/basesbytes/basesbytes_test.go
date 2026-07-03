package basesbytes

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestExample(t *testing.T) {
	var i int = 10
	var i64 int64 = 10
	var flag bool = true
	var letter rune = 'Ж'

	want := fmt.Sprintf("same value: 42 42 42 42\ndec=255 bin=11111111 oct=377 hex=ff\nsizeof int: %d\nsizeof int64: %d\nsizeof bool: %d\nsizeof rune: %d\ndec=64 bin=1000000 hex=40\n",
		unsafe.Sizeof(i),
		unsafe.Sizeof(i64),
		unsafe.Sizeof(flag),
		unsafe.Sizeof(letter),
	)

	if got := Example(); got != want {
		t.Fatalf("Example() = %q, want %q", got, want)
	}
}
