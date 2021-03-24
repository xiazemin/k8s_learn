FROM arm64v8/golang AS build-env
ENV GO111MODULE=on
ENV CGO_ENABLED 0
ADD  ./main.go /go/src/app/
ADD  ./go.mod /go/src/app/
WORKDIR /go/src/app/
RUN go build -v -o /go/src/app/server main.go
 
FROM scratch
COPY --from=build-env /go/src/app/server /app/server
WORKDIR /app
EXPOSE 5678
CMD [":5678"]
ENTRYPOINT ["./server"]