package numbers

import (
	"fmt"
	"strings"
)

func Example() string {
	var out strings.Builder

	// int имеет ширину 32 или 64 бита в зависимости от целевой архитектуры.
	var users int = 125
	fmt.Fprintln(&out, "users:", users)

	// Есть типы с точным размером: int8, int16, int32, int64.
	var small int8 = 120
	var big int64 = 9_000_000_000
	fmt.Fprintln(&out, "small/big:", small, big)

	// uint — только неотрицательные числа.
	var attempts uint = 3
	fmt.Fprintln(&out, "attempts:", attempts)

	// Арифметика.
	a, b := 17, 5
	fmt.Fprintln(&out, "a+b:", a+b)     // 22
	fmt.Fprintln(&out, "a-b:", a-b)     // 12
	fmt.Fprintln(&out, "a*b:", a*b)     // 85
	fmt.Fprintln(&out, "a/b:", a/b)     // 3, потому что оба числа int
	fmt.Fprintln(&out, "a mod b:", a%b) // 2, остаток от деления

	price := 1200
	discount := 150
	fmt.Fprintln(&out, "final price:", price-discount)

	return out.String()
}
