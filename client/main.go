package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
    commspb "github.com/kwdowicz/go-ms/comms"
)

func main() {
    // Connect to the gRPC server
	conn, err := grpc.NewClient("server:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    // Create a client for the PersonService
    client := commspb.NewPersonServiceClient(conn)

    // Create a new Person request
    person := &commspb.Person{
        Name: "Alice",
        Age:  30,
    }

    // Call the GreetPerson method
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    res, err := client.GreetPerson(ctx, person)
    if err != nil {
        log.Fatalf("Error calling GreetPerson: %v", err)
    }

    // Print the response
    log.Printf("Greeting: %s", res.Message)
}