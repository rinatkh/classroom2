# Live coding занятия 2

Основной маршрут — 90 минут, практика внутри тем. Все команды выполняются из корня classroom2. Каждый опыт занимает 2–5 минут: четыре изменения делаем вместе, оставшиеся можно разобрать готовыми или вынести в расширенный маршрут. Номера возврата к слайдам указаны в Speaker Notes и отдельной речи преподавателя.

## 0–10 минут: declaration

Открыть `cmd/01_declaration/main.go`, затем `internal/declaration/declaration.go`.

Команда: `make run-declaration`.

Ожидаемые строки: after +=: 15; after ++: 16.

Опыт: points=10 заменить на 20: получить 25 и 26. Вернуть 10.

Вопрос: Какая переменная новая в width, depth :=?

## 10–25 минут: numbers

Открыть `cmd/02_numbers/main.go`, затем `internal/numbers/numbers.go`.

Команда: `make run-numbers`.

Ожидаемые строки: a/b: 3; a mod b: 2.

Опыт: a=17 заменить на -17: получить -3 и -2. Вернуть 17.

Вопрос: Куда округляется целочисленное деление?

## 10–25 минут: bases-bytes

Открыть `cmd/03_bases_bytes/main.go`, затем `internal/basesbytes/basesbytes.go`.

Команда: `make run-bases-bytes`.

Ожидаемые строки: same value: 42 42 42 42; dec=255 bin=11111111 oct=377 hex=ff.

Опыт: x=64 заменить на 10: dec=10 bin=1010 hex=a. Вернуть 64.

Вопрос: Меняет ли %x само значение?

## 25–35 минут: float-bool-complex

Открыть `cmd/04_float_bool_complex/main.go`, затем `internal/floatboolcomplex/floatboolcomplex.go`.

Команда: `make run-float-bool-complex`.

Ожидаемые строки: final=239.88; can enter: true.

Опыт: age=20 заменить на 17: can enter: false. Вернуть 20.

Вопрос: && требует оба условия или одно?

## 35–50 минут: strings-runes

Открыть `cmd/05_strings_runes/main.go`, затем `internal/stringsrunes/stringsrunes.go`.

Команда: `make run-strings-runes`.

Ожидаемые строки: bytes: 16; runes: 10; first rune: П.

Опыт: text заменить на Go: обе длины 2. Вернуть Привет, Go.

Вопрос: Как получились 6×2+4 байта?

## 50–60 минут: constants-iota

Открыть `cmd/06_constants_iota/main.go`, затем `internal/constantsiota/constantsiota.go`.

Команда: `make run-constants-iota`.

Ожидаемые строки: statuses: 0 1 2; status text: paid.

Опыт: Добавить statusArchived. До нового case StatusText(3) даёт unknown; добавить case и тест archived=3. Прежнюю проверку неизвестного 3 заменить на 4. После опыта вернуть исходник и тесты.

Вопрос: Почему неизвестный статус всё равно возможен?

## 60–70 минут: conversion-named-types

Открыть `cmd/07_conversion_named_types/main.go`, затем `internal/conversionnamedtypes/conversionnamedtypes.go`.

Команда: `make run-conversion-named-types`.

Ожидаемые строки: sum: 30; atoi: 42 <nil>; invalid input: abc.

Опыт: В последнем Atoi заменить abc на 42: invalid input исчезнет. Вернуть abc. Дополнительно заменить 125 на abc: result не печатается.

Вопрос: Почему нельзя использовать число при err != nil?

## 70–85 минут: pointers

Открыть `cmd/08_pointers/main.go`, затем `internal/pointers/pointers.go`.

Команда: `make run-pointers`.

Ожидаемые строки: before: 10; after pointer change: 20; created with new: 42; nil pointer: true; default value: 100.

Опыт: *pointer=20 заменить на 30: value тоже 30. Вернуть 20 перед тестами.

Вопрос: Что копируется при передаче *int?

## 85–90 минут: тесты и homework

`make test-unit`, `make test-integration`, `make ci`. Unit вызывает Example, integration запускает реальный CLI. В красном тесте читаем имя, expected и actual, затем выясняем причину.

Домашняя работа: BASIC — 24 функции и две связанные группы констант; остальные 99 функций дополнительные. Starter компилируется, его функциональные тесты пока красные. `make ci-basic` принимает обязательную часть, `make ci` — полный банк.

## Расширенный маршрут

Комплексные числа — короткий обзор. Дополнительно: отрицательное деление; опыт 0.1+0.2 через переменные float64; составная Unicode-последовательность; функция с int против функции с *int; намеренная panic при разыменовании nil. После опыта с panic снова отключить строку. Не включать эти изменения в итоговый classroom.
