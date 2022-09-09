#!/bin/bash

# Set the correct arch
export GOOS=windows
export GOARCH=amd64

# Clean existing build
rm -f ./bin/double-tap.exe

# Create new build and setup .app structure
go build -o bin/double-tap.exe ./src

# Patch exe with correct icon
node ./rcedit.js

