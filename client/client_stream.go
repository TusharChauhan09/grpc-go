package main

import (
	"context"
	"log"
	"time"

	pb "github.com/TusharChauhan09/grpc-go/proto"
)

func callSayHelloClientStreaming(client pb.GreetServiceClient, names *pb.NamesList){
	log.Printf("client streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err!=nil{
		log.Fatalf("could not send names %v: ",err)
	}
	for _,name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err !=nil{
			log.Fatalf("error while client streaming: %v",err)
		}
		log.Printf("send the request with name: %v",name)
		time.Sleep(2*time.Second)
	}

	res,err := stream.CloseAndRecv()
	log.Printf("client streaming: %v ",res)
	if err != nil{
		log.Fatalf("error while recwving : %v",err)
	}
	log.Printf("%v", res.Messages)
}