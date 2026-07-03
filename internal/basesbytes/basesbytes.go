package basesbytes

import (
	"fmt"
	"strings"
	"unsafe"
)

func Example() string {
	var out strings.Builder

	// Числа можно записывать в разных системах счисления.
	decimal := 42      // десятичная
	hex := 0x2A        // шестнадцатеричная
	binary := 0b101010 // двоичная
	octal := 0o52      // восьмеричная

	fmt.Fprintln(&out, "same value:", decimal, hex, binary, octal) // 42 42 42 42

	// Вывод числа в разных системах счисления.
	n := uint64(255)
	fmt.Fprintf(&out, "dec=%d bin=%b oct=%o hex=%x\n", n, n, n, n)
	// dec=255 bin=11111111 oct=377 hex=ff

	// Сколько байтов занимает значение в памяти.
	var i int = 10
	var i64 int64 = 10
	var flag bool = true
	var letter rune = 'Ж'

	fmt.Fprintln(&out, "sizeof int:", unsafe.Sizeof(i))       // обычно 8 на 64-bit системе
	fmt.Fprintln(&out, "sizeof int64:", unsafe.Sizeof(i64))   // 8
	fmt.Fprintln(&out, "sizeof bool:", unsafe.Sizeof(flag))   // 1
	fmt.Fprintln(&out, "sizeof rune:", unsafe.Sizeof(letter)) // 4, rune == int32

	// Практика: выведите число 64 в десятичном, двоичном и шестнадцатеричном виде.
	// Ожидаемый вывод: dec=64 bin=1000000 hex=40
	//
	// Ответ:
	// x := 64
	// fmt.Printf("dec=%d bin=%b hex=%x\n", x, x, x)
	x := 64
	fmt.Fprintf(&out, "dec=%d bin=%b hex=%x\n", x, x, x)

	return out.String()
}
