#!/usr/bin/env bash

set -e

# get calling Directory
ORIG_DIR="$(pwd)"
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd "$SCRIPT_DIR"
cd ../

BANG="!!!!!!!!!!!!!!!!"

log() {
    echo "$BANG $1 $BANG"
    echo ""
}

while [[ "$#" -gt 0 ]]; do
    case $1 in
        --skip-ifaces) SKIP_GEN_IFACES=true ;;
        --skip-mocks) SKIP_GEN_MOCKS=true ;;
        --skip-tests) SKIP_GEN_TESTS=true ;;
        *) echo "Unknown parameter passed: $1"; exit 1 ;;
    esac
    shift
done

rm -rf service tests

if [ -z "$SKIP_GEN_IFACES" ]; then
    ./scripts/gen_ifaces.sh
    log DONE: gen_ifaces.sh
fi

log "Installing Go Dependencies to generate mocks"
npm run install:go

if [ -z "$SKIP_GEN_MOCKS"]; then
    ./scripts/gen_mocks.sh
    log "DONE: gen_mocks.sh"
fi

log "Installing Go Dependencies again for tests"
npm run install:go

if [ -z "$SKIP_GEN_TESTS" ]; then
    ./scripts/gen_tests.sh
    log "DONE: gen_tests.sh"
fi

cd "$ORIG_DIR"
