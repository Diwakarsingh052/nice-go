package main

import (
	"fmt"
	"log"
	pb "protoc-basics/proto"
)

func main() {
	simpleMessage()
}

func simpleMessage() {
	r := pb.BlogRequest{
		BlogId:  101,
		Title:   "Introduction to Protocol Buffers",
		Content: "Test",
	}

	fmt.Println(r.GetBlogId(), r.GetContent())

}

func nestedMessage() {

	// Create a new SearchResponse message with some results
	searchResponse := &pb.SearchResponse{
		Results: []*pb.SearchResponse_Result{
			{
				Url:   "https://grpc.io/",
				Title: "gRPC",
			},
			{
				Url:   "https://pkg.go.dev/",
				Title: "go packages",
			},
		},
	}
	log.Println("using nested types")
	// Print the URLs and titles of the results
	for _, result := range searchResponse.GetResults() {
		// adding result value in search result struct which was using
		fmt.Printf("Result URL: %s\n", result.GetUrl())
		fmt.Printf("Result Title: %s\n", result.GetTitle())
	}
}
