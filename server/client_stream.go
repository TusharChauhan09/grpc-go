package main

import (
	"io"
	"log"

	pb "github.com/TusharChauhan09/grpc-go/proto"
)

func (h *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF{
			return stream.SendAndClose(&pb.MessageList{Messages: messages})
		}
		if err !=nil {
			return err
		}
		log.Printf("got request with name: %v",req.Name)
		messages = append(messages,"hello "+req.Name)
	} 
}