package stringsrunes

import (
	"fmt"
	"strings"
	"unicode"
)

func Example() string {
	var out strings.Builder

	text := "Привет, Go" // 6 × 2 байта UTF-8 + 4 × 1 байт ASCII = 16 байт.
	// len для string считает байты, а не буквы.
	fmt.Fprintln(&out, "bytes:", len(text)) // кириллица занимает больше 1 байта 16

	// []rune декодирует кодовые точки Unicode; они не всегда совпадают с графемами.
	runes := []rune(text)
	fmt.Fprintln(&out, "runes:", len(runes))            // 10
	fmt.Fprintln(&out, "first rune:", string(runes[0])) // П

	// byte — это uint8. Удобен для ASCII/байтов.
	ascii := "Go"
	fmt.Fprintln(&out, "first byte:", ascii[0])                 // 71
	fmt.Fprintln(&out, "first byte as char:", string(ascii[0])) // G

	// Полезные функции strings.
	name := "  maria  " // "Maria"
	// В этом фиксированном примере имя после TrimSpace заведомо непустое.
	letters := []rune(strings.ToLower(strings.TrimSpace(name)))
	letters[0] = unicode.ToUpper(letters[0])
	normalized := string(letters)

	fmt.Fprintln(&out, "normalized:", normalized) // Maria

	s := "  golang  "
	fmt.Fprintln(&out, "upper:", strings.ToUpper(strings.TrimSpace(s))) // GOLANG

	return out.String()
}
