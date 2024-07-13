#!/usr/bin/env bash

set -e

# get calling Directory
ORIG_DIR="$(pwd)"
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

export PATH=$PATH:$(npm bin)

cd "$SCRIPT_DIR"
cd ../

AWS_SDK_VERSION=$(grep -E "aws/aws-sdk-go-v2 v[0-9]+\.[0-9]+\.[0-9]+" go.mod | cut -d' ' -f2)

DEST_DIR="tests/service"

help() {
    echo "Usage: $0 [-v|--version <AWS_SDK_VERSION>]"
}

while [[ "$#" -gt 0 ]]; do
    case $1 in
        -h|--help) help; exit 0 ;;
        -e|--extract-methods) service_dir="$2"; shift 1;;
        *) echo "Unknown parameter passed: $1"; exit 1 ;;
    esac
    shift
done

# rm -rf "$DEST_DIR" # refresh
# Ensure the destination directory exists
mkdir -p "$DEST_DIR"

echo "Generating Tests for AWS SDK Go v2 services Version $AWS_SDK_VERSION"

extract_test_methods() {
    local iface_file="$1"
    # shellcheck disable=SC2155
    local service_name=$(echo "$iface_file" | cut -d'/' -f3)

    # echo iface_file: "$iface_file"
    # echo service_name: "$service_name"
    # exit 0;

    # echo "Extracting test methods from $iface_file..."
    # Extract the method names from the interface file
    # and generate test methods for each
    # grep methods from iface.go file from type interface which should be NameFunc(....)

    # example:
    # type IClient interface {
    #     CreateBucket(input *CreateBucketInput) (*CreateBucketOutput, error)
    #     DeleteBucket(input *DeleteBucketInput) (*DeleteBucketOutput, error)
    #     ListBuckets(input *ListBucketsInput) (*ListBucketsOutput, error)
    # }

    grep -E '[A-Z][a-zA-Z0-9]*\(' "$iface_file" | cut -d'(' -f1 | while read -r line; do
        local method_name=$(echo "$line")
        if [ "$method_name" == "Options" ]; then
            # create Mock for go iface function Options() Options
            echo "
    t.Run(\"Test${method_name}\", func(t *testing.T) {
        output := ${service_name}.${method_name}{}
        mockClient.On(\"${method_name}\").Return(output)

        result := mockClient.${method_name}()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })"
            continue
        fi
            echo "
    t.Run(\"Test${method_name}\", func(t *testing.T) {
        input := &${service_name}.${method_name}Input{}
        output := &${service_name}.${method_name}Output{}

        mockClient.On(\"${method_name}\", ctx, input).Return(output, nil)

        result, err := mockClient.${method_name}(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })"
    done
}

if [ -n "$service_dir" ]; then
    extract_test_methods "$service_dir"
    exit 0
fi

rm -rf "$DEST_DIR" # refresh
mkdir -p "$DEST_DIR"
# Iterate over each service directory
for service_dir in ./service/*/*/iface.go; do
    # and not the "internal" directory
    if [ -f "$service_dir"  ] && [ "$(basename "$service_dir")" != "internal" ]; then
        service_name=$(echo "$service_dir" | cut -d'/' -f3)
        service_name_pascal=$(echo "$service_name" | pascalcase)
        
        test_service_dir="$DEST_DIR/${service_name}"
        mkdir -p "$test_service_dir"
        iface_file="$test_service_dir/iface_test.go"
        
        echo "Generating iface_test for $service_name..."

        # Generate the interface file
        # echo "$(pwd)/$service_dir"

# Fill out Go iface.go template
# TODO: figure out how to parllelize this
cat << EOF > "$iface_file"
package ${service_name}_test

// tests for the ${service_name} service interface for this ../../../service/${service_name}/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/${service_name}"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/${service_name}/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/${service_name}/${service_name}_iface"
	"github.com/stretchr/testify/assert"
)

func Test${service_name_pascal}ServiceCanBeMocked(t *testing.T) {
	var iface ${service_name}_iface.IClient
	iface = &${service_name}.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()
$(extract_test_methods "$service_dir")
}
EOF
    fi
done


echo "Test generation for interfaces and mocks complete."

cd "$ORIG_DIR"
