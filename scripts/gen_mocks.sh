#!/usr/bin/env bash

set -e

# get calling Directory
ORIG_DIR="$(pwd)"
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd "$SCRIPT_DIR"
cd ../

# Iterate over each local ./service directory
# generate mocks in parallel in batches of 15
i=0
for service_dir in ./service/*; do
    # and not the "internal" directory
    if [ -d "$service_dir"  ] && [ "$(basename "$service_dir")" != "internal" ]; then
        service_name=$(basename "$service_dir")
        echo "==============================="
        echo "Generating IClient mock for $service_name..."
        # execute in background bumping up against the limit of 5
        # capture the process id
        mockery --dir "$service_dir/${service_name}_iface" --name IClient --output "$service_dir/mocks" > /dev/null 2>&1 &
        # capture the process id
        pids[${i}]=$!
        i=$((i+1))
        if [ $i -eq 15 ]; then
            echo "Waiting for mocks to generate..."
            echo "!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!"
            # wait for all 5 to finish
            for pid in ${pids[*]}; do
                wait $pid
            done
            i=0
        fi
    fi
done

# wait for any remaining pids to finish
for pid in ${pids[*]}; do
    wait $pid
done
echo "Mocks generation complete."

cd "$ORIG_DIR"
