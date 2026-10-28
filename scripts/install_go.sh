#!/usr/bin/env bash

set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
cd "$SCRIPT_DIR/.."

SDK_SERVICE_PREFIX="github.com/aws/aws-sdk-go-v2/service/"

# -e keeps tidy going past modules it cannot resolve, and hides why.
# AWS tags hundreds of modules per release and the module proxy can fail
# on a just-tagged one for hours, so tidy silently drops that service.
go mod tidy -e

unresolved=$(go list -e -deps -f '{{if .Error}}{{.ImportPath}}{{end}}' ./... \
    | grep "^$SDK_SERVICE_PREFIX" \
    | cut -d/ -f1-5 \
    | sort -u || true)

for module in $unresolved; do
    echo "module proxy could not resolve $module, fetching from VCS"
    GOPROXY=direct GONOPROXY= GOFLAGS=-mod=mod go get "$module@latest"
done

if [ -n "$unresolved" ]; then
    go mod tidy
fi

go mod vendor
