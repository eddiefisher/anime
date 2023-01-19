.PHONY: build run help
.DEFAULT_GOAL := build

PROJECTNAME := $(shell basename "$(PWD)")

## build: build project
build:
	go build -o ./build/$(PROJECTNAME) -v ./cmd/

## run: run project
run:
	ENVIRONMENT="dev" \
	./build/$(PROJECTNAME)

## packages_update: recursively update Go packages
packages_update:
	go get -u ./... && go mod tidy

help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
