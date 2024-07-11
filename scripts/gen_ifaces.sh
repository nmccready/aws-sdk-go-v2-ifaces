#!/usr/bin/env bash

set -e

# get calling Directory
ORIG_DIR="$(pwd)"
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd "$SCRIPT_DIR"
cd ../

# ADD NPM BIN TO PATH
export PATH=$PATH:$(npm bin)
# Path to the AWS SDK Go v2 services directory
AWS_SDK_PATH="tmp/aws-sdk-go-v2/service"

# Destination for generated interface files
DEST_DIR="service"

rm -rf "$DEST_DIR" # refresh
# Ensure the destination directory exists
mkdir -p "$DEST_DIR"

# Function to extract public method signatures for Client struct receivers
extract_methods() {
    local service_dir="$1"
    # Extract public method signatures for Client struct receivers
    # e.g. func (c *Client) CreateBucket(input *CreateBucketInput) (*CreateBucketOutput, error)
    # sed removes the receiver and the method body, keep signature
    # cut removes remove leading bracket to method body
    # final cut removes the "$service_dir"/*.go file path and keeps signature
    grep -E 'func \(([^)]+)\*?Client\) [A-Z][a-zA-Z0-9]*\(' "$service_dir"/*.go| \
    sed -E 's/func \(([^)]+)\*?Client\)//'| \
    cut -d'{' -f1 | \
    cut -d':' -f2
}

# Iterate over each service directory
for service_dir in "$AWS_SDK_PATH"/*; do
    # and not the "internal" directory
    if [ -d "$service_dir"  ] && [ "$(basename "$service_dir")" != "internal" ]; then
        service_name=$(basename "$service_dir")
        mkdir -p "$DEST_DIR/${service_name}"
        iface_file="$DEST_DIR/${service_name}/iface.go"
        
        echo "Generating interface IClient for $service_name..."

        # Generate the interface file
        # echo "$(pwd)/$service_dir"

# Fill out Go iface.go template
cat << EOF > "$iface_file"

package $service_name

import (
    "github.com/aws/aws-sdk-go-v2/service/$service_name"
)

// IClient defines the interface for $service_name
type IClient interface {
$(extract_methods "$(pwd)/$service_dir")
}
EOF
    fi
done

echo "Interface generation complete."

cd "$ORIG_DIR"
