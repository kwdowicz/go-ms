package main

import (
	"context"
	_ "fmt"
	"log"
	"math/rand"
	"time"

	commspb "github.com/kwdowicz/go-ms/comms"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func createMsg(content string, tag string) *commspb.Msg {
    msg := &commspb.Msg{
		Content: content, 
        Tag: &commspb.Tag{Name: tag},
    }
    return msg
}

func sendMsg(client commspb.MsgServiceClient, ctx context.Context) *commspb.Ack {
    res, err := client.Post(ctx, createMsg("test content", "test tag")) 
    if err != nil {
        log.Fatalf("Error calling Post: %v", err)
    }
    return res
}

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
    // above if run locally and below if from docker-compose
	// conn, err := grpc.NewClient("server:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    client := commspb.NewMsgServiceClient(conn)
    ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
    defer cancel()

    for ;; {
        time.Sleep(time.Second * time.Duration(rand.Intn(5)))
        res := sendMsg(client, ctx)
        log.Printf("%s\n", res.Message)
    }
}