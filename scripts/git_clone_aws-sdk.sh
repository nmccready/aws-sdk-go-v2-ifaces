#!/usr/bin/env bash

set -e
# get calling Directory
ORIG_DIR="$(pwd)"
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd "$SCRIPT_DIR"
cd ../

mkdir -p ./tmp

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

# get the latest tag period
# AWS SDK V2 has changed to almost daily releases with no real root verion
# the typical release looks like "release-YYYY-MM-DD", which encapsulates a bunch of package releases.
# to see some grouped releases use:
#   git for-each-ref --sort=-creatordate --format '%(refname:short)' refs/tags | head -n 20
MOST_RECENT_TAG=$(git for-each-ref --sort=-creatordate --format '%(refname:short)' refs/tags | head -n 1)
# explicit checkout of the tagged version
git checkout "$MOST_RECENT_TAG"

cd "$ORIG_DIR"

# wipe out dependencies as they will come directly from pulling
# tmp/aws-sdk-go-v2 (structs, apis, etc)
rm -rf go.sum go.mod vendor

echo "module github.com/nmccready/aws-sdk-go-v2-ifaces

go 1.24" > go.mod
