package main

import (
	"io"
	"log"
	"time"

	pb "github.com/TusharChauhan09/grpc-go/proto"
)

func (h *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("received: %v", req.Name)

		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}

		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
}