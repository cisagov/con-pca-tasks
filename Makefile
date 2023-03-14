.PHONY: help run test update tidy version
include .env
export

# make all - Default Target. Does nothing.
all:
	@echo "Helper commands."
	@echo "For more information try 'make help'."

# target: help - Display callable targets.
help:
	@egrep "^# target:" [Mm]akefile

# target: run - run the application
run:
	go run *.go

# target: test - run application tests
test:
	go test -v ./...

# target: update - update all packages
update:
	go get all

# target: tidy - add missing necessary modules and remove unused modules
tidy:
	go mod tidy

# target: version - show application version
version:
	go run *.go -version
