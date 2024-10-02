# Makefile

.PHONY: all proto server client install

# The Go bin directory where 'go install' places binaries
GOBIN := $(shell go env GOPATH)/bin

# Ensure GOBIN is in PATH
export PATH := $(GOBIN):$(PATH)

# Proto source and destination
PROTO_SRC := ./comms/comms.proto
PROTO_DEST := ./comms

# Default target
all: install proto

# Install required tools
install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Compile proto files
proto:
	protoc -I=./comms --go_out=$(PROTO_DEST) --go-grpc_out=$(PROTO_DEST) $(PROTO_SRC)

# Run the server
server:
	go run ./server/main.go

# Run the client
client:
	go run ./client/main.go