#!/bin/bash

# Set the correct arch
export GOOS=darwin
export GOARCH=amd64

# Clean existing build
rm -rf ./bin/DoubleTap.app

# Create new build and setup .app structure
go build -o bin/double-tap ./src
cp -R ./resources/darwin/DoubleTap.app ./bin/
mkdir ./bin/DoubleTap.app/Contents/MacOS
cp ./bin/double-tap ./bin/DoubleTap.app/Contents/MacOS/

# Remove the leftover binary
rm ./bin/double-tap
