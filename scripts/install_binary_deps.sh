#!/usr/bin/env bash

if ! command -v mockery >/dev/null 2>&1 || [[ "$version" != "$(mockery --version 2>/dev/null | tail -n 1 | grep -E 'v\d+\.\d+\.\d+')" ]]; then
    go install github.com/vektra/mockery/v2@v2.53.4
fi

if ! command -v golangci-lint >/dev/null 2>&1 || [[ "$(golangci-lint version | cut -d ' ' -f 4)" != "v1.64.6" ]]; then
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.6
fi
