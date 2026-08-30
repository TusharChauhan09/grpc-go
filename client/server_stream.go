package main

import (
	"context"
	"io"
	"log"

	pb "github.com/TusharChauhan09/grpc-go/proto"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient,names *pb.NamesList) {
	log.Printf("streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err!=nil{
		log.Printf("error while making server streaming call: %v",err)
	}
	for {
		message , err := stream.Recv()
		if err == io.EOF{
			break
		}
		log.Printf("message recieved: %v: ", message)
	}

}