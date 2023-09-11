package main

import (
	pb "client/gen/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"sync"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)

	}
	defer conn.Close()

	if err != nil {
		log.Fatalln(err)

	}
	client := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	stream, err := client.GreetEveryone(ctx)
	if err != nil {
		log.Fatalf("failed to call GreetEveryone stream: %v\n", err)
	}

	requests := []*pb.GreetEveryoneRequest{
		{FirstName: "John"},
		{FirstName: "Bruce"},
		{FirstName: "Roy"},
	}
	//using WaitGroup to wait for goroutines to finish
	wg := &sync.WaitGroup{}

	wg.Add(2)
	// First goroutine to handle send operation // send request to the remote service
	go func() {
		defer wg.Done()
		for _, req := range requests {
			log.Printf("Sending message: %v\n", req)
			//sending requests
			err := stream.Send(req)

			if err != nil {
				if closeErr := stream.CloseSend(); closeErr != nil {
					log.Printf("Failed to close stream: %v", closeErr)
					return
				}
				return
			}
			time.Sleep(1 * time.Second)

		}
		//closing stream when the server finished sending
		err := stream.CloseSend()
		if err != nil {
			log.Println(err)
			return
		}
	}()

	//recv the values from the remote service
	go func() {
		defer wg.Done()
		for {

			res, err := stream.Recv()
			if err == io.EOF {
				log.Printf("stream has ended")
				break
			}
			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}
			select {
			case <-stream.Context().Done():
				log.Println("remote service cancelled the request")
				return
			default:
				// The remote service is still connected
			}

			log.Printf("Received: %v\n", res.Result)

		}

	}()

	wg.Wait()
	fmt.Println("end of bidirectional communication")

}
