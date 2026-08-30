package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/TusharChauhan09/grpc-go/proto"
)

func callSayHelloBidirectionalStreaming (client pb.GreetServiceClient, names *pb.NamesList){
	log.Printf("bi directional stream start")
	stream,err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil{
		 log.Fatalf("error while sending names: %v",err)
	}

	waitc := make(chan struct{})
	go func (){
		for{

			message, err := stream.Recv()
			if err==io.EOF{
				break 
			}
			if err!=nil{
				log.Fatalf("error while streaming data : %v",err)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req:= &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err!=nil{
			log.Fatalf("error while sending req: %v",err);
		}
		time.Sleep(2*time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("bidir streaming finished")

}