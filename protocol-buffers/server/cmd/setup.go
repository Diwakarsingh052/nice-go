package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "server/gen/proto"
)

func main() {
	listener, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Println(err)
		return
	}

	//NewServer creates a gRPC server which has no service registered
	//and has not started to accept requests yet.
	s := grpc.NewServer()

	pb.RegisterUserServiceServer(s, &userService{})

	//Serve accepts incoming connections on the listener lis
	err = s.Serve(listener)
	if err != nil {
		fmt.Println(err)
		return
	}
}
