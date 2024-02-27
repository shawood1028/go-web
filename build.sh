#!/usr/bin/env bash
# linux amd64 打包
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main
ssh -i /Users/ohla/Desktop/shawood_alicloud.pem root@47.243.133.65 "rm -f /data/code/shawood-site/bin/main"
scp -i /Users/ohla/Desktop/shawood_alicloud.pem ./bin/main root@47.243.133.65:/data/code/shawood-site/bin/
scp -r -i /Users/ohla/Desktop/shawood_alicloud.pem ./etc root@47.243.133.65:/data/code/shawood-site/
ssh -i /Users/ohla/Desktop/shawood_alicloud.pem root@47.243.133.65 "supervisorctl restart shawood-site"