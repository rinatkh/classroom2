package declaration

import (
	"fmt"
	"strings"
)

func Example() string {
	var out strings.Builder

	// 1. Объявление переменной без значения: получим zero value.
	var count int
	fmt.Fprintln(&out, "zero int:", count) // 0

	// 2. Объявление с явным типом.
	var lessonNumber int = 2
	fmt.Fprintln(&out, "lesson:", lessonNumber) // 2

	// 3. Go может сам вывести тип по значению.
	var language = "Go"
	fmt.Fprintln(&out, "language:", language) // Go

	// 4. Короткое объявление работает только внутри функций.
	studentName := "Maria"
	fmt.Fprintln(&out, "student:", studentName) // Maria

	// 5. += и ++
	points := 10
	points += 5                             // points = points + 5
	fmt.Fprintln(&out, "after +=:", points) // 15

	points++                                // points = points + 1 , points += 1
	fmt.Fprintln(&out, "after ++:", points) // 16

	// В Go нет ++points и нет --points. Только points++ / points-- отдельной строкой.

	// 6. Несколько переменных сразу.
	var width, height int = 10, 20
	fmt.Fprintln(&out, "rectangle:", width, height) // 10 20

	// 7. Можно переприсваивать сразу несколько значений.
	width, height = 15, 25
	fmt.Fprintln(&out, "new rectangle:", width, height) // 15 25

	// 8. В коротком объявлении хотя бы одна переменная должна быть новой.
	width, depth := 30, 5
	fmt.Fprintln(&out, "box:", width, height, depth) // 30 25 5

	name := "Maria"
	age := 25
	city := "Berlin"
	fmt.Fprintln(&out, "profile:", name, age, city)

	return out.String()
}
