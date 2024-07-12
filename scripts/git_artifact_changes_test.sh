#!/usr/bin/env bash

set -e

# get calling Directory
ORIG_DIR="$(pwd)"
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd "$SCRIPT_DIR"
cd ../

# use git and exit with an error if there are pending changes
if ! git diff --quiet; then
    echo "There are pending changes in the working directory."
    echo "Failing the test. As running npm run sync should not have changes."
    exit 1
fi

cd "$ORIG_DIR"
