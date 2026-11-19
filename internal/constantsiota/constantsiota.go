package constantsiota

import (
	"fmt"
	"strings"
)

const (
	statusNew = iota
	statusPaid
	statusCanceled
)

// StatusText возвращает текст известного статуса и обрабатывает неизвестное значение.
func StatusText(status int) string {
	switch status {
	case statusNew:
		return "new"
	case statusPaid:
		return "paid"
	case statusCanceled:
		return "canceled"
	default:
		return "unknown"
	}
}

func Example() string {
	var out strings.Builder

	const appName = "go-course"
	const maxAttempts = 3

	fmt.Fprintln(&out, "app:", appName)              // go-course
	fmt.Fprintln(&out, "max attempts:", maxAttempts) // 3

	// iota удобно использовать для enum-like наборов значений.
	fmt.Fprintln(&out, "statuses:", statusNew, statusPaid, statusCanceled) // 0 1 2

	status := statusPaid
	fmt.Fprintln(&out, "is paid:", status == statusPaid) // true
	fmt.Fprintln(&out, "status text:", StatusText(status))

	const (
		roleUser = iota
		roleManager
		roleAdmin
	)
	fmt.Fprintln(&out, "roles:", roleUser, roleManager, roleAdmin)

	return out.String()
}
