run:
	go run cmd/main.go

get:
	go get ./...

fmt:
	go fmt ./...

lint:
	go vet ./...

build:
	go build -o app cmd/main.go

release:
	go build -o app -ldflags "-s -w" cmd/main.go