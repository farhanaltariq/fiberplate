.PHONY: run test

run:
	swag init && go run .
build:
	go build -o=./bin/app .
test:
	go test -v --cover ./...
