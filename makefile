.PHONY: run test

run:
	go run .
build:
	go build -o=./bin/app .
test:
	go test -v --cover ./...
