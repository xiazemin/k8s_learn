% go mod init banana
go: creating new go.mod: module banana

docker buildx build \
  --platform linux/arm64 -t banana:v1 .


docker  build  -t banana:v1 .
docker build -t apple:v1 .