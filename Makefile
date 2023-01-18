.PHONY: help run test tidy
include .env

# target: help - display callable targets.
help:
	@egrep "^# target:" [Mm]akefile

# target: run - run the application
run:
	go run *.go

# target: test - run application tests
test:
	go test -v ./...

# target: tidy - add missing necessary modules and remove unused modules
tidy:
	go mod tidy
