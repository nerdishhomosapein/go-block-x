install :
	go install all

watch :
	~/go/bin/air

build:
	go build -o ./bin/go-block-x

run: build
	./bin/go-block-x

lint :
	go fmt ./... &&  golangci-lint -D errcheck run

test :
	~/go/bin/gotestsum --format testname --junitfile unit-tests.xml -- -race -covermode=atomic -coverprofile=coverage.out ./...

html-coverage-report :
	go tool cover -html=coverage.out
