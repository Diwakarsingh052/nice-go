package main

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	pb "server/gen/proto"
)

func (us *userService) GreetEveryone(stream pb.UserService_GreetEveryoneServer) error {
	log.Println("GreetEveryone was invoked")

	for {
		//receiving the streaming request from the client
		req, err := stream.Recv()

		//If the client has finished sending the request, we will quit
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Printf("error while reading client stream: %v\n", err)
			return status.Error(codes.Internal, err.Error())
		}

		select {
		case <-stream.Context().Done():
			log.Println("remote service cancelled the request")
			return status.Error(codes.Unavailable, "remote service disconnected")
		default:
			// The remote service is still connected
		}
		// write domain logic that work on the basis of the req received
		res := "Hello " + req.FirstName + "!"

		//send the response
		err = stream.Send(&pb.GreetEveryoneResponse{
			Result: res,
		})

		if err != nil {
			log.Printf("error while sending data to client: %v\n", err)
			return status.Error(codes.Internal, err.Error())
		}

	}
}
