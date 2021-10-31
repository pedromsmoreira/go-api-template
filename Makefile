.PHONY: build run test unit-tests

unit-tests:
	go test -v ./...

run: 
	go run cmd/main.go

build:
	docker build -t go-api-template:local .

test:
	docker run --rm -it -p 3000:3000 go-api-template:local