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

	err := session.UserSession().Set(r.UserSessionRedisDB, tokenPayload, uint(req.UserID), req.Roles)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}

	err = session.UserSession().Set(r.UserSessionRedisDB, tokenPayload2, uint(req.UserID), req.Roles)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}

	access, err := util.Jwt().New(r.Env.Secret, tokenPayload, req.Roles, time.Minute * time.Duration(req.AccessTokenExpireInMinutes))
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}

	refresh, err := util.Jwt().New(r.Env.Secret, tokenPayload2, req.Roles, time.Minute * time.Duration(req.RefreshTokenExpireInMinutes))
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}
	return &pb.SetUserSessionResponse{
		Token:   access,
		Refresh: refresh,
	}, nil
}

func (r *SessionService) DeleteUserSession(ctx context.Context, req *pb.DeleteUserSessionPayload) (*pb.OkResponse, error) {
	err := session.UserSession().DelUserSessionBySessionID(r.UserSessionRedisDB, req.Token)
	if err != nil {
		return nil, err
	}
	return &pb.OkResponse{
		Ok: true,
	}, nil
}

func (r *SessionService) ClearUserAllSession(ctx context.Context, req *pb.ClearUserAllSessionPayload) (*pb.OkResponse, error) {
	err := session.UserSession().DelUserSessionByUserID(r.UserSessionRedisDB, uint(req.UserID))
	if err != nil {
		return nil, err
	}
	return &pb.OkResponse{
		Ok: true,
	}, nil
}

func (r *SessionService) GetUserSessionRefreshToken(ctx context.Context, req *pb.GetUserSessionRefreshTokenPayload) (*pb.SetUserSessionResponse, error) {
	key, _, err := util.Jwt().ExtractClaims(r.Env.Secret, req.RefreshToken)
	if err != nil {
		return nil, err
	}
	
	sessionData, err := session.UserSession().GetOneBySessionID(r.UserSessionRedisDB, key)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}

	tokenPayload := util.NewUUID()
	
	err = session.UserSession().Set(r.UserSessionRedisDB, tokenPayload, sessionData.UserID, sessionData.Roles)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}
	
	access, err := util.Jwt().New(r.Env.Secret, tokenPayload, sessionData.Roles, time.Minute * time.Duration(req.AccessTokenExpireInMinutes))
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}
	
	return &pb.SetUserSessionResponse{
		Token: access,
		Refresh: req.RefreshToken,
	}, nil
}
func (r *SessionService) VerifyUserSession(ctx context.Context, req *pb.VerifyUserSessionParams) (*pb.VerifyUserSessionResponse, error) {
	key, _, err := util.Jwt().ExtractClaims(r.Env.Secret, req.Token)
	if err != nil {
		return nil, err
	}

	data, err := session.UserSession().GetOneBySessionID(r.UserSessionRedisDB, key)
	if err != nil {
		return nil, err
	}

	return &pb.VerifyUserSessionResponse{
		Ok:     true,
		UserID: uint64(data.UserID),
		Roles:  data.Roles,
	}, nil
}
