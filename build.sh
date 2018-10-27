#!/bin/bash
echo "Building project..."
rm -rf output && 
mkdir output &&
go build -o output/secret main.go &&
cp -v output/secret /Users/stefanlapcevic/go/bin
echo "Done"