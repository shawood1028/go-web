#!/usr/bin/env bash
# linux amd64 打包
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main
scp -i /Users/ohla/Desktop/shawood_alicloud.pem ./bin/main root@47.243.133.65:/data/code/bin/
scp -r -i /Users/ohla/Desktop/shawood_alicloud.pem ./etc root@47.243.133.65:/data/code/