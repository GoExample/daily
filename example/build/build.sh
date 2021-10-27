#!/usr/bin/env bash

flags="-s -w -X 'main.buildStamp=$(date -u '+%Y-%m-%d %I:%M:%S')' -X 'main.gitHash=$(git describe --long --dirty --abbrev=14)' -X 'main.goVersion=$(go version)'"

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "${flags}"