package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	commspb "github.com/kwdowicz/go-ms/comms"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func createMsg() *commspb.Msg {
    rand.Seed(time.Now().UnixMilli()) // deprecated, maybe remove it completly

    subjects := []string{"The cat", "A dog", "My neighbor", "Tosia", "Gabi", "Staś"}
	verbs := []string{"eats", "runs", "jumps", "sleeps", "talks", "cries"}
	objects := []string{"pizza", "over the fence", "at the park", "with joy", "loudly", "like a pro"}

    sentence := fmt.Sprintf("%s %s %s.", randomWord(subjects), randomWord(verbs), randomWord(objects))

    msg := &commspb.Msg{
		Content: sentence, 
    }
    return msg
}

func randomWord(words []string) string {
    return words[rand.Intn(len(words))]
}

func sendMsg(client commspb.PersonServiceClient, ctx context.Context) *commspb.Ack {
    res, err := client.Post(ctx, createMsg())
    if err != nil {
        log.Fatalf("Error calling Post: %v", err)
    }
    return res
}

func main() {
	// conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
    // above if run locally and below if from docker-compose
	conn, err := grpc.NewClient("server:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    client := commspb.NewPersonServiceClient(conn)
    ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
    defer cancel()

    for ;; {
        time.Sleep(time.Second * time.Duration(rand.Intn(5)))
        res := sendMsg(client, ctx)
        log.Printf("%s\n", res.Message)
    }
}