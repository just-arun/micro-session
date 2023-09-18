package rpcservice

import (
	"context"
	"fmt"

	pb "github.com/just-arun/micro-session-proto"
)

func (r *SessionService) SetUserSession(ctx context.Context, req *pb.UserSessionPayload) (*pb.SetUserSessionResponse, error) {
	fmt.Printf("DATA RECEIVED \nROLE: %v\nUSER_ID:%v\n", req.Role, req.UserID)
	return &pb.SetUserSessionResponse{
		Token: "some-random-token-shit",
	}, nil
}
