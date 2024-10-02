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
    commspb.UnimplementedMsgServiceServer
}

func (s *Server) Post(ctx context.Context, msg *commspb.Msg) (*commspb.Ack, error) {
    log.Printf("Client sent: %v, %v", msg.Content, msg.Tag)
    message := fmt.Sprintf("Server recieved this from you: %v in tag: %v", msg.Content, msg.Tag)
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
    commspb.RegisterMsgServiceServer(grpcServer, &Server{})

    log.Println("gRPC server is running on port 50051...")

    // Start the server
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}