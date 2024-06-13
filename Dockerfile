# Build Golang binary
FROM golang:1.21.3 AS build-golang

WORKDIR /app/go-rpc

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download
# Copy the source code
COPY . .
RUN go build cmd/infofile/server.go
COPY . .

RUN mkdir /data
COPY ./resources/*.* /data/

VOLUME /data
EXPOSE 50051


CMD ["./server"]
