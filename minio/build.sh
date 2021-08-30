go mod init minio
go mod tidy
tar -czvf test.tar ./* 
go run main.go 
