package main

import (
	pb "client/gen/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func main() {
	dialOpts := []grpc.DialOption{
		//interceptors
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:5001", dialOpts...)
	defer conn.Close()
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	//creates a gRPC client for the UserService service defined in the pb package
	//and binds it to the specified connection conn.
	client := pb.NewUserServiceClient(conn)

	//constructing SignupRequest
	req := &pb.SignupRequest{
		User: &pb.User{
			Name:     "John",
			Email:    "john@email.com",
			Password: "abc",
			Roles:    []string{"ADMIN", "USER"},
		},
	}

	//creating context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	res, err := client.Signup(ctx, req)

	if err != nil {
		grpcStatus, ok := status.FromError(err)
		if ok {
			log.Println(grpcStatus.Code())
			log.Println(grpcStatus.Message())
			return
		}
		log.Fatalf("failed: %v", err)
	}

	log.Println(res.Result)

}
