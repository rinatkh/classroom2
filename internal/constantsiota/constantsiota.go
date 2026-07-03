package constantsiota

import (
	"fmt"
	"strings"
)

func Example() string {
	var out strings.Builder

	const appName = "go-course"
	const maxAttempts = 3

	const (
		statusNew = iota
		statusPaid
		statusCanceled
	)

	fmt.Fprintln(&out, "app:", appName)
	fmt.Fprintln(&out, "max attempts:", maxAttempts)

	// iota удобно использовать для enum-like наборов значений.
	fmt.Fprintln(&out, "statuses:", statusNew, statusPaid, statusCanceled) // 0 1 2

	status := statusPaid
	fmt.Fprintln(&out, "is paid:", status == statusPaid) // true

	// Практика: добавьте константы ролей RoleUser, RoleManager, RoleAdmin через iota.
	// Выведите их значения.
	// Ожидаемый вывод: roles: 0 1 2
	//
	// Ответ:
	// const (
	//     RoleUser = iota
	//     RoleManager
	//     RoleAdmin
	// )
	// fmt.Println("roles:", RoleUser, RoleManager, RoleAdmin)
	const (
		roleUser = iota
		roleManager
		roleAdmin
	)
	fmt.Fprintln(&out, "roles:", roleUser, roleManager, roleAdmin)

	return out.String()
}
