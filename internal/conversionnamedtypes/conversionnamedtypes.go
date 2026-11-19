package conversionnamedtypes

import (
	"fmt"
	"strconv"
	"strings"
)

func Example() string {
	var out strings.Builder

	type userID int64
	type rub int
	type kop int

	// Go не делает неявные преобразования между числовыми типами.
	var a int = 10
	var b int64 = 20
	fmt.Fprintln(&out, "sum:", int64(a)+b) // 30

	// Пользовательские типы помогают выражать смысл.
	var id userID = 1001
	fmt.Fprintln(&out, "user id:", id)

	// Деньги лучше хранить в минимальных единицах: копейки, центы.
	var price rub = 1200
	priceInKop := kop(price * 100)
	fmt.Fprintln(&out, "kop:", priceInKop)

	// strconv для преобразований string <-> number.
	n, err := strconv.Atoi("42")
	if err != nil {
		fmt.Fprintln(&out, "error:", n, err) // 0 error
	} else {
		fmt.Fprintln(&out, "atoi:", n, err) // 42 <nil>
		str := strconv.Itoa(n)
		fmt.Fprintln(&out, "itoa:", str) // "42"
	}

	value, parseErr := strconv.Atoi("125")
	if parseErr != nil {
		fmt.Fprintln(&out, "error:", value, parseErr) // 0 error
	} else {
		fmt.Fprintln(&out, "result:", value+5)
	}
	fmt.Fprintln(&out, "parse error:", parseErr)

	// При ошибке не используем результат парсинга в дальнейших вычислениях.
	_, invalidErr := strconv.Atoi("abc")
	if invalidErr != nil {
		fmt.Fprintln(&out, "invalid input:", "abc")
	}

	return out.String()
}
