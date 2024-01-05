FROM golang:alpine AS builder
ENV CGO_ENABLED 0
ENV GOOS linux
RUN apk update --no-cache && \
    apk add curl && \
    apk add unzip
WORKDIR /build
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
ENV PROTOC_ZIP=protoc-3.15.8-linux-x86_64.zip
RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.15.8/$PROTOC_ZIP && \
    unzip -o $PROTOC_ZIP -d ~/protoc