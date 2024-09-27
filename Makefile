# Variables
PROTO_DIR := comms
PROTO_FILES := $(PROTO_DIR)/*.proto
GO_OUT_DIR := $(PROTO_DIR)
GO_FILES := $(wildcard $(GO_OUT_DIR)/*.pb.go)

# Protobuf compilation
.PHONY: proto
proto: $(PROTO_FILES)
	protoc --go_out=$(GO_OUT_DIR) --go-grpc_out=$(GO_OUT_DIR) --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative $(PROTO_FILES)

# Run the server
.PHONY: run-server
run-server: proto
	go run server/main.go

# Run the client
.PHONY: run-client
run-client: proto
	go run client/main.go

# Clean generated protobuf files
.PHONY: clean
clean:
	rm -f $(GO_FILES)