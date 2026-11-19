SHELL := /bin/bash

GO ?= go
GO_PACKAGES := ./...
UNIT_PACKAGES := ./internal/...
INTEGRATION_PACKAGES := ./test/integration/...
BIN_DIR := bin
COVERAGE_FILE ?= coverage.out
COVERAGE_THRESHOLD ?= 90.0
PACKAGE_FILE := $(BIN_DIR)/classroom2-linux-amd64.tar.gz

.PHONY: \
	help \
	run-declaration \
	run-numbers \
	run-bases-bytes \
	run-float-bool-complex \
	run-strings-runes \
	run-constants-iota \
	run-conversion-named-types \
	run-pointers \
	run-all \
	build \
	build-declaration \
	build-numbers \
	build-bases-bytes \
	build-float-bool-complex \
	build-strings-runes \
	build-constants-iota \
	build-conversion-named-types \
	build-pointers \
	package \
	deps-check \
	mod-check \
	test \
	test-unit \
	test-integration \
	test-race \
	coverage \
	coverage-check \
	fmt \
	fmt-check \
	vet \
	check \
	ci \
	clean

help:
	@echo "Available commands:"
	@echo "  make run-declaration             - run declaration example"
	@echo "  make run-numbers                 - run integer numbers example"
	@echo "  make run-bases-bytes             - run bases, bytes and memory size example"
	@echo "  make run-float-bool-complex      - run float, bool and complex example"
	@echo "  make run-strings-runes           - run strings, bytes and runes example"
	@echo "  make run-constants-iota          - run constants and iota example"
	@echo "  make run-conversion-named-types  - run conversion and named types example"
	@echo "  make run-pointers                - run pointers example"
	@echo "  make run-all                     - run all examples"
	@echo "  make build                       - build all binaries into bin/"
	@echo "  make package                     - build all binaries and create tar.gz artifact"
	@echo "  make deps-check                  - download and verify Go modules"
	@echo "  make mod-check                   - check that go.mod/go.sum are tidy"
	@echo "  make test                        - run unit and integration tests"
	@echo "  make test-unit                   - run unit tests for internal packages"
	@echo "  make test-integration            - run integration tests for CLI output"
	@echo "  make test-race                   - run unit tests with race detector"
	@echo "  make coverage                    - run unit coverage and print report"
	@echo "  make coverage-check              - fail if unit coverage is below threshold"
	@echo "  make fmt                         - format Go files"
	@echo "  make fmt-check                   - check formatting without changing files"
	@echo "  make vet                         - run go vet"
	@echo "  make check                       - run fast local checks"
	@echo "  make ci                          - run full CI checks locally"
	@echo "  make clean                       - remove generated artifacts"

run-declaration:
	$(GO) run ./cmd/01_declaration

run-numbers:
	$(GO) run ./cmd/02_numbers

run-bases-bytes:
	$(GO) run ./cmd/03_bases_bytes

run-float-bool-complex:
	$(GO) run ./cmd/04_float_bool_complex

run-strings-runes:
	$(GO) run ./cmd/05_strings_runes

run-constants-iota:
	$(GO) run ./cmd/06_constants_iota

run-conversion-named-types:
	$(GO) run ./cmd/07_conversion_named_types

run-pointers:
	$(GO) run ./cmd/08_pointers

run-all: \
	run-declaration \
	run-numbers \
	run-bases-bytes \
	run-float-bool-complex \
	run-strings-runes \
	run-constants-iota \
	run-conversion-named-types \
	run-pointers

build: \
	build-declaration \
	build-numbers \
	build-bases-bytes \
	build-float-bool-complex \
	build-strings-runes \
	build-constants-iota \
	build-conversion-named-types \
	build-pointers

build-declaration:
	@mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN_DIR)/01_declaration ./cmd/01_declaration

build-numbers:
	@mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN_DIR)/02_numbers ./cmd/02_numbers

build-bases-bytes:
	@mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN_DIR)/03_bases_bytes ./cmd/03_bases_bytes

build-float-bool-complex:
	@mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN_DIR)/04_float_bool_complex ./cmd/04_float_bool_complex

build-strings-runes:
	@mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN_DIR)/05_strings_runes ./cmd/05_strings_runes

build-constants-iota:
	@mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN_DIR)/06_constants_iota ./cmd/06_constants_iota

build-conversion-named-types:
	@mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN_DIR)/07_conversion_named_types ./cmd/07_conversion_named_types

build-pointers:
	@mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN_DIR)/08_pointers ./cmd/08_pointers

package:
	$(MAKE) clean
	$(MAKE) build GOOS=linux GOARCH=amd64
	tar -czf $(PACKAGE_FILE) -C $(BIN_DIR) \
		01_declaration \
		02_numbers \
		03_bases_bytes \
		04_float_bool_complex \
		05_strings_runes \
		06_constants_iota \
		07_conversion_named_types \
		08_pointers

deps-check:
	$(GO) mod download
	$(GO) mod verify

mod-check:
	$(GO) mod tidy
	@if git rev-parse --is-inside-work-tree >/dev/null 2>&1; then \
		git diff --exit-code -- go.mod go.sum; \
	else \
		echo "Skipping git diff because this directory is not a git repository"; \
	fi

test: test-unit test-integration

test-unit:
	$(GO) test $(UNIT_PACKAGES)

test-integration:
	$(GO) test $(INTEGRATION_PACKAGES)

test-race:
	$(GO) test -race $(UNIT_PACKAGES)

coverage:
	$(GO) test $(UNIT_PACKAGES) -covermode=atomic -coverprofile=$(COVERAGE_FILE)
	$(GO) tool cover -func=$(COVERAGE_FILE)

coverage-check: coverage
	@coverage="$$( $(GO) tool cover -func=$(COVERAGE_FILE) | awk '/^total:/ {gsub("%", "", $$3); print $$3}')"; \
	awk -v coverage="$$coverage" -v threshold="$(COVERAGE_THRESHOLD)" 'BEGIN { \
		if (coverage + 0 < threshold + 0) { \
			printf "coverage %.1f%% is below threshold %.1f%%\n", coverage, threshold; \
			exit 1; \
		} \
		printf "coverage %.1f%% is enough; threshold %.1f%%\n", coverage, threshold; \
	}'

fmt:
	gofmt -w $$(find . -name '*.go' -not -path './$(BIN_DIR)/*')

fmt-check:
	@files="$$(gofmt -l $$(find . -name '*.go' -not -path './$(BIN_DIR)/*'))"; \
	if [ -n "$$files" ]; then \
		echo "Go files are not formatted:"; \
		echo "$$files"; \
		exit 1; \
	fi

vet:
	$(GO) vet $(GO_PACKAGES)

check: fmt-check vet test

ci: deps-check mod-check fmt-check vet test test-race coverage-check build

clean:
	rm -rf $(BIN_DIR) $(COVERAGE_FILE)
