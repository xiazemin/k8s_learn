FROM golang:alpine AS build-env
ENV GO111MODULE=on
ENV CGO_ENABLED 0
ADD  ./server/main.go /go/src/app/
ADD  ./server/go.mod /go/src/app/
WORKDIR /go/src/app/
RUN go build -v -o /go/src/app/server main.go
 
FROM arm64v8/centos:7
COPY --from=build-env /go/src/app/server /app/server
WORKDIR /app
EXPOSE 10248
CMD [":10248"]
ENTRYPOINT ["./server"]