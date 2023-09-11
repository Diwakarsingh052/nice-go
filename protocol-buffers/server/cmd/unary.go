package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	pb "server/gen/proto"
)

type userService struct {
	pb.UnimplementedUserServiceServer
}
type User struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	Roles        []string `json:"roles"`
	PasswordHash string   `json:"-"`
}

func (us userService) Signup(ctx context.Context, req *pb.SignupRequest) (*pb.SignupResponse, error) {
	fmt.Println(req.GetUser())
	err := errors.New("values missing")
	nu := req.GetUser() // fetching the request sent by the client
	if nu.Name == "" {
		return nil, status.Errorf(codes.Internal, "account creation failed %v", err)
	}
	u := User{
		ID:    101,
		Name:  nu.Name,
		Email: nu.Email,
		Roles: nu.Roles,
	}
	//write queries to put this user in db

	log.Println(u)
	return &pb.SignupResponse{Result: u.Email + " account created"}, nil

}
