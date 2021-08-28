go mod init redis
go mod tidy
go run redis-cli.go 
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o redis-cli redis-cli.go