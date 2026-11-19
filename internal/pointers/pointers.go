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
	*pointer = 20
	fmt.Fprintln(&out, "after pointer change:", value) // 20

	// new(T) создает значение типа T и возвращает указатель на него.
	created := new(int) // Ненулевой *int указывает на нулевое int-значение.
	*created = 42
	fmt.Fprintln(&out, "created with new:", *created) // 42

	// nil pointer — указатель, который никуда не указывает.
	var missing *int // nil
	// *missing = 42 // Намеренно отключено: разыменование nil вызывает panic.
	fmt.Fprintln(&out, "nil pointer:", missing == nil) // true

	// Перед разыменованием указателя обычно нужно проверить его на nil.
	defaultValue := 100
	if missing == nil {
		fmt.Fprintln(&out, "default value:", defaultValue)
	}

	return out.String()
}
