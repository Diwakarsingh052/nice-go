package main

import (
	pb "client/gen/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
)

//In server streaming, the server sends back a sequence of responses
//after getting the client’s request message.

func main() {
	conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)

	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	req := pb.GetPostsRequest{UserId: 101}

	stream, err := client.GetPosts(context.Background(), &req)

	if err != nil {
		log.Fatalln(err)
	}
	for {
		//receiving values from stream
		post, err := stream.Recv() // blocking call // it would wait until new messages are not sent
		//if the server has finished sending the request, we will quit
		if err == io.EOF {
			break
		}
		//any other kind of error would be caught here
		if err != nil {
			log.Println(err)
			return
		}
		select {
		case <-stream.Context().Done():
			log.Println("remote service cancelled the request")
			return
		default:
			// The Client is still connected
		}

		fmt.Println("reading stream")
		//printing data received
		fmt.Println(post)
		fmt.Println()

	}

}
