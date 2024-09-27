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

// GreetPerson implements the GreetPerson RPC method
func (s *Server) GreetPerson(ctx context.Context, person *commspb.Person) (*commspb.Greeting, error) {
    // Create a greeting message
    message := fmt.Sprintf("Hello, %s! You are %d years old.", person.Name, person.Age)
    return &commspb.Greeting{Message: message}, nil
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