.PHONY: run test fmt vet bench clean

DAY ?= day91

run:
	go run ./$(DAY)

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

bench:
	go test -bench=. ./...

clean:
	go clean -testcache
