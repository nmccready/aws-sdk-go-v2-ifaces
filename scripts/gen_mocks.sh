#!/usr/bin/env bash

set -e

# get calling Directory
ORIG_DIR="$(pwd)"
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd "$SCRIPT_DIR"
cd ../

# Parse arguments for sync flag
SYNC_MODE=0
for arg in "$@"; do
    if [ "$arg" == "--sync" ]; then
        SYNC_MODE=1
    fi
done

# Iterate over each local ./service directory
# generate mocks in parallel in batches of 15
# If SYNC_MODE=1, run synchronously
i=0
for service_dir in ./service/*; do
    # and not the "internal" directory
    if [ -d "$service_dir"  ] && [ "$(basename "$service_dir")" != "internal" ]; then
        service_name=$(basename "$service_dir")
        echo "==============================="
        echo "Generating IClient mock for $service_name..."
        if [ $SYNC_MODE -eq 1 ]; then
            # Synchronous mode
            mockery --dir "$service_dir/${service_name}_iface" --name IClient --output "$service_dir/mocks"
        else
            # Async mode (default)
            mockery --dir "$service_dir/${service_name}_iface" --name IClient --output "$service_dir/mocks" > /dev/null 2>&1 &
            pids[${i}]=$!
            i=$((i+1))
            if [ $i -eq 15 ]; then
                echo "Waiting for mocks to generate..."
                echo "!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!"
                for pid in ${pids[*]}; do
                    wait $pid
                done
                i=0
            fi
        fi
    fi
done

# wait for any remaining pids to finish (only if async)
if [ $SYNC_MODE -eq 0 ]; then
    for pid in ${pids[*]}; do
        wait $pid
    done
fi
echo "Mocks generation complete."

cd "$ORIG_DIR"
