#!/usr/bin/env bash

set -ex
cd "$(dirname "${BASH_SOURCE[0]}")"

CMD_GOLINT="go run golang.org/x/lint/golint@latest"
CMD_STATICCHECK="go run honnef.co/go/tools/cmd/staticcheck@latest"
CMD_GOCOV="go run github.com/axw/gocov/gocov@latest"
CMD_GOCOV_HTML="go run github.com/matm/gocov-html/cmd/gocov-html@latest "

DIR_BUILD=".build"
DIR_COVERAGE="$DIR_BUILD/coverage"

rm -rf "$DIR_BUILD"
mkdir -p "$DIR_COVERAGE"

go mod tidy
go generate ./...
go build -v ./...
diff -u <(echo -n) <(gofmt -d ./)
$CMD_GOLINT -set_exit_status ./...
go vet ./...
$CMD_STATICCHECK ./...
go test -trimpath -race -failfast -shuffle=on -covermode=atomic -coverprofile="$DIR_COVERAGE/coverage.out" -count=1 ./...
$CMD_GOCOV convert "$DIR_COVERAGE/coverage.out" > "$DIR_COVERAGE/coverage.json"
$CMD_GOCOV_HTML -t golang < "$DIR_COVERAGE/coverage.json"> "$DIR_COVERAGE/coverage.html"
