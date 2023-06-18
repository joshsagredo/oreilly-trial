#!/bin/bash

go test -tags "unit e2e integration" ./... -race -p 1 -coverprofile=coverage.txt -covermode=atomic -ldflags="-X github.com/bilalcaliskan/oreilly-trial/internal/mail.token=${API_TOKEN}"
go tool cover -html=all_coverage.txt -o all_cover.html
open all_cover.html
