package stringsrunes

import (
	"fmt"
	"strings"
)

func Example() string {
	var out strings.Builder

	text := "Привет, Go"

	// len для string считает байты, а не буквы.
	fmt.Fprintln(&out, "bytes:", len(text)) // кириллица занимает больше 1 байта

	// []rune позволяет работать с символами Unicode.
	runes := []rune(text)
	fmt.Fprintln(&out, "runes:", len(runes))
	fmt.Fprintln(&out, "first rune:", string(runes[0])) // П

	// byte — это uint8. Удобен для ASCII/байтов.
	ascii := "Go"
	fmt.Fprintln(&out, "first byte:", ascii[0])                 // 71
	fmt.Fprintln(&out, "first byte as char:", string(ascii[0])) // G

	// Полезные функции strings.
	name := "  maria  "
	normalized := strings.Title(strings.TrimSpace(name))
	fmt.Fprintln(&out, "normalized:", normalized) // Maria

	// Практика: дана строка "  golang  ". Уберите пробелы и переведите в верхний регистр.
	// Ожидаемый вывод: GOLANG
	//
	// Ответ:
	// s := "  golang  "
	// fmt.Println(strings.ToUpper(strings.TrimSpace(s)))
	s := "  golang  "
	fmt.Fprintln(&out, "upper:", strings.ToUpper(strings.TrimSpace(s)))

	return out.String()
}
