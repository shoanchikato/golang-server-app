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

tidy:
	go mod tidy

t:
	go test ./test/...

tv:
	go test -count=1 -v ./test/...

build:
	go build -o app cmd/server/*.go

release:
	go build -o app -ldflags "-s -w" cmd/server/*.go