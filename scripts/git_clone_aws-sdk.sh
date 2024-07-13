#!/usr/bin/env bash

set -e
# get calling Directory
ORIG_DIR="$(pwd)"
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd "$SCRIPT_DIR"
cd ../

mkdir -p ./tmp

# cut the version out of the go.mod file, get the version of the aws-sdk-go-v2
AWS_SDK_VERSION=$(grep -E "aws/aws-sdk-go-v2 v[0-9]+\.[0-9]+\.[0-9]+" go.mod | cut -d' ' -f2)
cd ./tmp

# clone only if aws-sdk-go-v2 is not already cloned
if [ -d "aws-sdk-go-v2" ]; then
    echo "AWS SDK Go v2 already cloned"
else
    echo "Cloning AWS SDK Go v2 version $AWS_SDK_VERSION"
    # Clone the AWS SDK Go v2 repository
    git clone https://github.com/aws/aws-sdk-go-v2.git
fi

# Checkout the specified version
cd aws-sdk-go-v2
cp ../../.prototools .
git checkout main
git pull -p
git checkout "$AWS_SDK_VERSION"

cd "$ORIG_DIR"

