set GOARCH=amd64
set GOOS=linux
go build .\src\main.go
scp -i C:\Users\shawood\Downloads\shawood_alicloud.pem ./main.exe root@47.243.133.65:~