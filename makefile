.PHONY: run build test lint clean

run:
	go run cmd/api/main.go

build:
	go build -o bin/api cmd/api/main.go

test:
	go test ./... -v -coverprofile=coverage.out

lint:
	golangci-lint run

lint-fix:
	gofmt -w .
	goimports -local github.com/AndrKuz/furutori -w .

clean:
	rm -rf bin/ coverage.out