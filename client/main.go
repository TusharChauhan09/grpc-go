package main

import (
	"log"

	pb "github.com/TusharChauhan09/grpc-go/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	//? grpc client connection to the server
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalf("did not connect to grpc server: %v",err)
	}
	defer conn.Close()
 
	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Alex", "Bob", "Charlie"},
	}

	callSayHello(client)
	callSayHelloServerStreaming(client,names)
	callSayHelloClientStreaming(client,names)
	callSayHelloBidirectionalStreaming(client,names)
}