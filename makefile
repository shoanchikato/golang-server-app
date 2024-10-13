run-r: doc build
	./bin/app

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
	go build -o ./bin/app cmd/server/*.go

release:
	go build -o ./bin/app -ldflags "-s -w" cmd/server/*.go

clean:
	rm -rf ./bin ./docs

doc: doc-fmt
	swag init -g cmd/server/main.go

doc-fmt:
	swag fmt -g cmd/server/main.go