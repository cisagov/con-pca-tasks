.PHONY: all help build logs loc up stop down
include .env

# target: run - run the application
run:
	go run *.go

# target: test - run application tests
test:
	go test
