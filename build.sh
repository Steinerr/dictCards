#!/usr/bin/env bash
#go get github.com/lib/pq
rm -rf bin/
mkdir bin/
cp -r app/templates bin/
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/wiki app/wiki.go