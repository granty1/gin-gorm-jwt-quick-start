.PHONY: all build test
run:
	@echo Start run...
	@go run .
build:
	@go mod tidy
	@go build .

