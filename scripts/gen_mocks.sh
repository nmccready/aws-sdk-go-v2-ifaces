#!/usr/bin/env bash

set -e

# get calling Directory
ORIG_DIR="$(pwd)"
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd "$SCRIPT_DIR"
cd ../

# Iterate over each local ./service directory
for service_dir in ./service/*; do
    # and not the "internal" directory
    if [ -d "$service_dir"  ] && [ "$(basename "$service_dir")" != "internal" ]; then
    mockery --dir "$service_dir" --name IClient --output "$service_dir/mocks"
    fi
done

echo "Mocks generation complete."

cd "$ORIG_DIR"
