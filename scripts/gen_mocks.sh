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

# Add arrays to track pids and failures, and helper to capture wait statuses
pids=()
failed_pids=()
failed_statuses=()

# create a temporary log directory for background jobs
LOG_DIR="${SCRIPT_DIR}/logs/mocks"
rm -rf "$LOG_DIR" 
mkdir -p "$LOG_DIR"

pid_log_map=()

wait_capture() {
    # Temporarily disable errexit so we can capture wait's exit code and show logs
    local pid status logfile
    for pid in "$@"; do
        if [ -z "$pid" ]; then
            continue
        fi
        set +e
        wait "$pid"
        status=$?
        set -e
        logfile="${pid_log_map[$pid]}"
        if [ $status -ne 0 ]; then
            echo "ERROR: mockery process (pid $pid) exited with status $status"
            if [ -n "$logfile" ] && [ -f "$logfile" ]; then
                echo "---- Begin log for pid $pid ($logfile) ----"
                tail -n 500 "$logfile"
                echo "---- End log for pid $pid ($logfile) ----"
            else
                echo "No logfile found for pid $pid"
            fi
            failed_pids+=("$pid")
            failed_statuses+=("$status")
        fi
    done
}

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
            # capture both stdout and stderr to a per-service logfile
            logfile="$LOG_DIR/${service_name}.$(date +%s%N).log"
            mockery --dir "$service_dir/${service_name}_iface" --name IClient --output "$service_dir/mocks" > "$logfile" 2>&1 &
            pid=$!
            pids[${i}]=$pid
            pid_log_map["$pid"]="$logfile"
            i=$((i+1))
            if [ $i -eq 15 ]; then
                echo "Waiting for mocks to generate..."
                echo "!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!"
                # capture wait errors without letting set -e abort the script
                wait_capture "${pids[@]}"
                # reset batch state
                pids=()
                i=0
            fi
        fi
    fi
done

# wait for any remaining pids to finish (only if async)
if [ $SYNC_MODE -eq 0 ]; then
    if [ ${#pids[@]} -ne 0 ]; then
        wait_capture "${pids[@]}"
    fi
fi

# If any background job failed, report and exit with the first non-zero status
if [ ${#failed_statuses[@]} -ne 0 ]; then
    echo "One or more mockery processes failed. PIDs: ${failed_pids[*]} Statuses: ${failed_statuses[*]}"
    echo "Logs are available in: $LOG_DIR"
    exit "${failed_statuses[0]}"
else
    # no failures — remove logs to keep workspace clean
    rm -rf "$LOG_DIR"
fi

echo "Mocks generation complete."

cd "$ORIG_DIR"
