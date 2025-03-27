FROM golang:1.22 as builder
RUN apt-get update && apt-get -y install make  
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.59.1
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
