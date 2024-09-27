package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    commspb "github.com/kwdowicz/go-ms/comms"
)

// Server is used to implement the gRPC PersonService server
type Server struct {
    commspb.UnimplementedPersonServiceServer
}

func (s *Server) Post(ctx context.Context, msg *commspb.Msg) (*commspb.Ack, error) {
    log.Printf("Client sent: %v", msg.Content)
    message := fmt.Sprintf("Server recieved this from you: %v", msg.Content)
    return &commspb.Ack{Message: message}, nil
}

func main() {
    // Start listening on port 50051
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    // Create a new gRPC server
    grpcServer := grpc.NewServer()

    // Register the PersonService server
    commspb.RegisterPersonServiceServer(grpcServer, &Server{})

    log.Println("gRPC server is running on port 50051...")

    // Start the server
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}