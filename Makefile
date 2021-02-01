test:
	go test .

build:
	go build -o bin/main *.go

run:
	go run *.go

compile:
	# 32-Bit Systems
	# FreeBDS
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 *.go
	# MacOS
	GOOS=darwin GOARCH=386 go build -o bin/main-darwin-386 *.go
	# Linux
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 *.go
	# Windows
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 *.go
        # 64-Bit
	# FreeBDS
	GOOS=freebsd GOARCH=amd64 go build -o bin/main-freebsd-amd64 *.go
	# MacOS
	GOOS=darwin GOARCH=amd64 go build -o bin/main-darwin-amd64 *.go
	# Linux
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-amd64 *.go
	# Windows
	GOOS=windows GOARCH=amd64 go build -o bin/main-windows-amd64 *.go


all: test run
