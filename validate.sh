#!/usr/bin/env bash

set -ex
cd "$(dirname "${BASH_SOURCE[0]}")"

CMD_GOLANGCI_LINT="go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.2"
CMD_GOCOV="go run github.com/axw/gocov/gocov@latest"
CMD_GOCOV_HTML="go run github.com/matm/gocov-html/cmd/gocov-html@latest "

DIR_BUILD=".build"
DIR_COVERAGE="$DIR_BUILD/coverage"

rm -rf "$DIR_BUILD"
mkdir -p "$DIR_COVERAGE"

go mod tidy
go generate ./...
go build -v ./...
$CMD_GOLANGCI_LINT run
go test -trimpath -race -failfast -shuffle=on -covermode=atomic -coverprofile="$DIR_COVERAGE/coverage.out" -count=1 ./...
$CMD_GOCOV convert "$DIR_COVERAGE/coverage.out" > "$DIR_COVERAGE/coverage.json"
$CMD_GOCOV_HTML -t golang < "$DIR_COVERAGE/coverage.json"> "$DIR_COVERAGE/coverage.html"
