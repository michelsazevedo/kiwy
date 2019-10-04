-include .env.dev

.PHONY: all

GOFILES := $(wildcard *.go)

run:
	go run main.go

test:
	go test $(GOFILES)
