#!/usr/bin/env bash

set -e

# get calling Directory
ORIG_DIR="$(pwd)"
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd "$SCRIPT_DIR"
cd ../

rm -rf service tests

./gen_ifaces.sh
./gen_mocks.sh
./gen_tests.sh

cd "$ORIG_DIR"
