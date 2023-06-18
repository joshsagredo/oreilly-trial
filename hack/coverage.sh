#!/bin/bash

go test -tags "unit e2e integration" ./... -race -p 1 -coverprofile=coverage.txt -covermode=atomic -ldflags="-X github.com/bilalcaliskan/oreilly-trial/internal/mail.token=${API_TOKEN}"
go tool cover -html=coverage.txt -o cover.html
open cover.html
