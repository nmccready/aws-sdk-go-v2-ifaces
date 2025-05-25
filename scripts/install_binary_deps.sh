#!/usr/bin/env bash

version=$(mockery --version 2>/dev/null | tail -n 1 | grep -oP 'v\d+\.\d+\.\d+')

if ! command -v mockery >/dev/null 2>&1 || [[ "$version" != "v2.53.4" ]]; then
    go install github.com/vektra/mockery/v2@v2.53.4
fi
