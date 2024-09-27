# Stage 1: Build the Go applications (server and client)
FROM golang:1.23-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the server
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./server

# Build the client
RUN CGO_ENABLED=0 GOOS=linux go build -o client ./client

# Stage 2: Run the server in a lightweight container
FROM alpine:latest AS server

# Set the working directory
WORKDIR /app

# Copy the server binary from the build stage
COPY --from=builder /app/server .

# Expose the port used by the gRPC server
EXPOSE 50051

# Command to run the gRPC server
CMD ["./server"]

# Stage 3: Run the client in a lightweight container
FROM alpine:latest AS client

# Set the working directory
WORKDIR /app

# Copy the client binary from the build stage
COPY --from=builder /app/client .

# Command to run the gRPC client
CMD ["./client"]