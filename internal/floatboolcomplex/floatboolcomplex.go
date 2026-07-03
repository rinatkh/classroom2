package floatboolcomplex

import (
	"fmt"
	"strings"
)

func Example() string {
	var out strings.Builder

	// float64 — основной тип для дробных чисел.
	price := 199.90
	tax := 0.20
	final := price * (1 + tax)
	fmt.Fprintf(&out, "final=%.2f\n", final) // final=239.88

	// Важно: float не подходит для точных денег в production.
	// Для обучения и простых расчетов можно использовать, но деньги обычно хранят в копейках/центах.

	// bool принимает только true/false.
	age := 20
	hasTicket := true
	canEnter := age >= 18 && hasTicket
	fmt.Fprintln(&out, "can enter:", canEnter) // true

	// Логические операторы: &&, ||, !
	isWeekend := false
	isHoliday := true
	canRest := isWeekend || isHoliday
	fmt.Fprintln(&out, "can rest:", canRest) // true
	fmt.Fprintln(&out, "workday:", !canRest) // false

	// complex встречается редко, но в Go есть встроенно.
	z := 2 + 3i
	fmt.Fprintln(&out, "complex:", z)

	// Практика: посчитайте среднюю оценку 4.5, 5.0, 3.5 и проверьте, что она >= 4.0.
	// Ожидаемый вывод:
	// average=4.33
	// passed=true
	//
	// Ответ:
	// avg := (4.5 + 5.0 + 3.5) / 3
	// fmt.Printf("average=%.2f\n", avg)
	// fmt.Println("passed=", avg >= 4.0)
	avg := (4.5 + 5.0 + 3.5) / 3
	fmt.Fprintf(&out, "average=%.2f\n", avg)
	fmt.Fprintln(&out, "passed:", avg >= 4.0)

	return out.String()
}
