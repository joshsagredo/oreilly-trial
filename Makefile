lint:
	golangci-lint run

fmt:
	go fmt ./...

vet:
	go vet ./...

ineffassign:
	go get github.com/gordonklaus/ineffassign
	go mod vendor
	ineffassign ./...

test:
	go test ./...

test_coverage:
	go test ./... -race -coverprofile=coverage.txt -covermode=atomic

build:
	go build -o bin/main main.go

run:
	go run main.go

cross-compile:
	# 32-Bit Systems
	# FreeBDS
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	# MacOS
	GOOS=darwin GOARCH=386 go build -o bin/main-darwin-386 main.go
	# Linux
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	# Windows
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go
        # 64-Bit
	# FreeBDS
	GOOS=freebsd GOARCH=amd64 go build -o bin/main-freebsd-amd64 main.go
	# MacOS
	GOOS=darwin GOARCH=amd64 go build -o bin/main-darwin-amd64 main.go
	# Linux
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-amd64 main.go
	# Windows
	GOOS=windows GOARCH=amd64 go build -o bin/main-windows-amd64 main.go

upgrade-direct-deps:
	for item in `grep -v 'indirect' go.mod | grep '/' | cut -d ' ' -f 1`; do \
		echo "trying to upgrade direct dependency $$item" ; \
		go get -u $$item ; \
  	done
	go mod tidy
	go mod vendor

all: test build run
