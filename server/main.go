package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/TusharChauhan09/grpc-go/proto"
)

const (
	port = ":8080"
)

//? to implemet grpc services we use this struct 
type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	//! tcp listener
	listener, err := net.Listen("tcp",port)
	if err!=nil{
		log.Fatalf("failed to start the server %v",err)
	}

	//! grpc service server
	grpcServer := grpc.NewServer()

	//? bind the service with the server
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("server started at %v",listener.Addr())

	//? grpc server serve over tcp port 8080
	if err := grpcServer.Serve(listener); err != nil{
		log.Fatalf("failed to start: %v",err)
	}
}