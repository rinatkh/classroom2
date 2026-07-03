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
	fmt.Fprintln(&out, "atoi:", n, err) // 42 <nil>

	str := strconv.Itoa(42)
	fmt.Fprintln(&out, "itoa:", str) // "42"

	// Практика: преобразуйте строку "125" в int, прибавьте 5 и выведите результат.
	// Ожидаемый вывод: result: 130
	//
	// Ответ:
	// value, err := strconv.Atoi("125")
	// if err != nil {
	//     fmt.Println("error:", err)
	//     return
	// }
	// fmt.Println("result:", value+5)
	value, parseErr := strconv.Atoi("125")
	fmt.Fprintln(&out, "result:", value+5)
	fmt.Fprintln(&out, "parse error:", parseErr)

	return out.String()
}
