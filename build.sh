#!/usr/bin/env bash
rm -rf bin/
mkdir bin/
cp -r app/templates bin/
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/wiki app/wiki.go