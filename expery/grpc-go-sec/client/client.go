package main

import (
	"context"
	"log"

	"grpc-go-sec/chat"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)
	bs := chat.NewBroadcastServiceClient(conn)

	response, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

	responseM, errM := bs.Greet(context.Background(), &chat.RequestMessage{RequestString: "airbender!"})

	if errM != nil {
		log.Fatalf("Error when calling Greet: %s", errM)
	}
	log.Printf("%s", responseM.ResponseString)
	responseM, errM = bs.SayPeace(context.Background(), &chat.RequestMessage{RequestString: "peace"})
	if errM != nil {
		log.Fatalf("Error when calling SayPeace: %s", errM)
	}
	log.Printf("%s", responseM.ResponseString)
}
