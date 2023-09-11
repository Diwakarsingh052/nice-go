package main

import (
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	pb "server/gen/proto"
	"sync"
	"time"
)

// In client-streaming RPC, the client sends multiple messages/request to the server
// instead of a single request.
// The server sends back a single response to the client.

func (us *userService) CreatePost(stream pb.UserService_CreatePostServer) error {
	wg := &sync.WaitGroup{}

	// Receive CreatePost request from a client in batches
	for {

		//receiving the request from the client
		req, err := stream.Recv()

		//If the client has finished sending the request, we will quit
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		//latency of server processing
		//in the meantime when the server was doing the processing of the request,
		//if the request is cancelled, then we would detect that using select
		time.Sleep(time.Second * 4)
		//during the request if a client close the connection we will know inside this select block

		select {
		//this case evaluates if a client is disconnected
		case <-stream.Context().Done():
			log.Println("client cancelled the request")
			return status.Error(codes.Internal, "client disconnected")
		default:
			// The Client is still connected move on for further processing
		}
		// Process creates post request
		b, _ := json.MarshalIndent(req, "", " ")
		log.Printf("Received Create Post Requests: %v", string(b))

		posts := req.GetPosts()
		log.Println("adding all the posts into the db")
		wg.Add(1)
		go AddPost(posts, wg)

	}
	return stream.SendAndClose(&pb.CreatePostResponse{Result: "done"})

}

func AddPost(posts []*pb.Post, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, p := range posts {
		time.Sleep(2 * time.Second)
		log.Println("adding post ", p.Title)
	}
}
