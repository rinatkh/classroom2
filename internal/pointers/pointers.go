package pointers

import (
	"fmt"
	"strings"
)

func Example() string {
	var out strings.Builder

	// Указатель хранит адрес переменной в памяти.
	value := 10
	pointer := &value
	fmt.Fprintln(&out, "before:", value) // 10

	// Через *pointer можно изменить исходное значение.
	(*pointer)++
	fmt.Fprintln(&out, "after pointer change:", value) // 11

	// new(T) создает значение типа T и возвращает указатель на него.
	created := new(int)
	*created = 42
	fmt.Fprintln(&out, "created with new:", *created) // 42

	// nil pointer — указатель, который никуда не указывает.
	var missing *int
	fmt.Fprintln(&out, "nil pointer:", missing == nil) // true

	// Перед разыменованием указателя обычно нужно проверить его на nil.
	defaultValue := 100
	if missing == nil {
		fmt.Fprintln(&out, "default value:", defaultValue)
	}

	return out.String()
}
