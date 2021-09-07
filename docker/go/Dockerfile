FROM arm64v8/golang:1.16

RUN mkdir -p /etc/apk

RUN echo "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.4/main" > /etc/apk/repositories

RUN apt install apk

RUN apk add --update curl bash
