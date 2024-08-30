run-r: build
	./app

run:
	go run cmd/server/*.go

get:
	go get ./...

fmt:
	go fmt ./...

lint:
	go vet --race ./...

build:
	go build -o app cmd/server/*.go

release:
	go build -o app -ldflags "-s -w" cmd/server/*.go