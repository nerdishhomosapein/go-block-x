build:
	go build -o ./bin/go-block-x

run: build
	./bin/go-block-x

test:
	go test -v ./...

lint:
	go fmt ./...
