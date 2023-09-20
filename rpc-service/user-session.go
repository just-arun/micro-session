package rpcservice

import (
	"context"
	"fmt"
	"time"

	pb "github.com/just-arun/micro-session-proto"
	"github.com/just-arun/micro-session/session"
	"github.com/just-arun/micro-session/util"
)

func (r *SessionService) SetUserSession(ctx context.Context, req *pb.UserSessionPayload) (*pb.SetUserSessionResponse, error) {
	fmt.Println(req)
	tokenPayload := util.NewUUID()
	tokenPayload2 := util.NewUUID()
	access, err := util.Jwt().New(r.Env.Secret, tokenPayload, time.Minute*10)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}
	session.UserSession().Set(r.UserSessionRedisDB, tokenPayload, uint(req.UserID), req.Role)
	session.UserSession().Set(r.UserSessionRedisDB, tokenPayload2, uint(req.UserID), req.Role)
	refresh, err := util.Jwt().New(r.Env.Secret, tokenPayload2, time.Hour*2)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}
	fmt.Println(access, refresh)
	return &pb.SetUserSessionResponse{
		Token:   access,
		Refresh: refresh,
	}, nil
}

