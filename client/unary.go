package main

import (
	"context"
	"log"
	"time"

	pb "github.com/TusharChauhan09/grpc-go/proto"
)

func callSayHello(client pb.GreetServiceClient){
	ctx , cancel := context.WithTimeout(context.Background(),time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil{
		log.Fatalf("failed to retrive respone on unary: %v",err)
	}

	log.Printf("reponse: %v",res.Message)
}