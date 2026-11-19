# Classroom 2: переменные, типы и память

Продолжение занятия 1: main вызывает небольшую функцию, команда печатает результат, тесты закрепляют поведение. Сохраняем восемь существующих cmd и internal packages. Настройку Go и Git заново не изучаем.

## Окружение и структура

Go 1.22+, Make, Git. Module: `github.com/rinatkh/classroom2`. Примеры совместимы с Go 1.22; версии 1.23+ специально не требуются.

`cmd/<номер_темы>/main.go` вызывает `Example()` из соответствующего `internal` package. Unit-тесты находятся рядом, CLI integration — в `test/integration/classroom_test.go`. Числовые размеры, вычисляемые через unsafe.Sizeof, зависят от целевой сборки.

## Запуск

Из корня classroom2:

```bash
make help
make run-all
make run-declaration
make run-numbers
make run-bases-bytes
make run-float-bool-complex
make run-strings-runes
make run-constants-iota
make run-conversion-named-types
make run-pointers
```

Каждый пример можно менять на 2–5 минут: предсказать результат, изменить одну строку, запустить, объяснить и восстановить исходное значение. Точный маршрут находится в [docs/live-coding-script.md](docs/live-coding-script.md). За 90 минут достаточно четырёх изменений руками, остальные команды разбираются готовыми; оставшиеся опыты переходят в расширенный маршрут.

## Проверки

```bash
make fmt-check
make vet
make test-unit
make test-integration
make test-race
go test -race ./...
make coverage-check
make ci
make build
make package
```

`make test` объединяет unit и integration. `make ci` дополнительно проверяет зависимости, tidy, форматирование, vet, race, покрытие не ниже 90% и сборку. `make package` собирает Linux/amd64 независимо от архитектуры компьютера; `make build` собирает под текущую. `make clean` удаляет артефакты. В Git-репозитории mod-check сравнивает go.mod/go.sum с HEAD, поэтому после применения патча сохраните изменения обычным commit перед полным CI.

Соседний homework_go_2 — самостоятельный module. Classroom его не импортирует.
