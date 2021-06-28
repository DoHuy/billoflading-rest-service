#!/bin/bash
# shellcheck disable=SC2164

go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/vektra/mockery/cmd/mockery

# install and scan lint in code golang
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.31.0
# install and scan security code golang
curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b $GOPATH/bin v2.4.0
