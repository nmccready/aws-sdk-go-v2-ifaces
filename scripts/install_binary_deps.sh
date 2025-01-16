#!/usr/bin/env bash

set -e

if ! mockery --version > /dev/null 2>&1; then
    go install github.com/vektra/mockery/v2@v2.43.2
fi
