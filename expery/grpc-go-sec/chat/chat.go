package chat

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedChatServiceServer
}

type BServer struct {
	UnimplementedBroadcastServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}

func (bs *Server) SayPeace(ctx context.Context, req *RequestMessage) (*ResponseMessage, error) {
	response := "and Love"
	if strings.ToLower(req.RequestString) != "peace" {
		response = "I don't know"
	}
	log.Printf("Received message body form client: %s", req.RequestString)
	return &ResponseMessage{ResponseString: response}, nil
}

func (bs *BServer) Greet(ctx context.Context, request *RequestMessage) (*ResponseMessage, error) {
	log.Printf("Received Request from client: %s", request.RequestString)
	responseInfo := fmt.Sprintf("Hello from the server to %s", request.RequestString)
	return &ResponseMessage{ResponseString: responseInfo}, nil
}
