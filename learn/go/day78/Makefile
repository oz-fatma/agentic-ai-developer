# Day 78: Makefile for learn/go module
# Run from learn/go: make -f day78/Makefile help

GO ?= go
DAYS := $(shell seq 61 80)

.PHONY: help fmt vet test run-all tidy

help:
	@echo "Targets: fmt vet test run-all tidy run-day"

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...

test:
	$(GO) test ./day66/... ./day68/... ./day70/...

run-all:
	@for d in $(DAYS); do \
		echo "==> day$$d"; \
		$(GO) run ./day$$d || exit 1; \
	done

run-day:
	@test -n "$(DAY)" || (echo "usage: make -f day78/Makefile run-day DAY=76" && exit 1)
	$(GO) run ./day$(DAY)

tidy:
	$(GO) mod tidy
