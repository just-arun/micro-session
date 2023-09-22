package rpcservice

import (
	"context"

	"github.com/just-arun/micro-session/model"
	"github.com/just-arun/micro-session/session"
	pb "github.com/just-arun/micro-session-proto"
)

func (r *SessionService) SetOTP(ctx context.Context, req *pb.OTPPayload) (*pb.OkResponse, error) {
	err := session.
		OTP().
		SetOTP(r.GeneralSessionRedisDB.Conn(), model.OTP{
			OTP:    req.Otp,
			Key:    req.Key,
			UserID: uint(req.UserID),
		}, r.Env.OTP.ExpireTime)
	if err != nil {
		return nil, err
	}
	return &pb.OkResponse{
		Ok: true,
	}, nil
}

func (r *SessionService) GetOTP(ctx context.Context, req *pb.OTPPayload) (*pb.OkResponse, error) {
	_, err := session.
		OTP().
		GetOTP(r.GeneralSessionRedisDB.Conn(),
			model.OTP{
				OTP:    req.Otp,
				Key:    req.Key,
				UserID: uint(req.UserID),
			}, r.Env.OTP.ExpireTime)
	if err != nil {
		return nil, err
	}
	return &pb.OkResponse{
		Ok: true,
	}, nil
}

func (r *SessionService) GetAndExpireOTP(ctx context.Context, req *pb.OTPPayload) (*pb.OkResponse, error) {
	_, err := session.
		OTP().
		GetAndDelOTP(r.GeneralSessionRedisDB.Conn(),
			model.OTP{
				OTP:    req.Otp,
				Key:    req.Key,
				UserID: uint(req.UserID),
			}, r.Env.OTP.ExpireTime)
	if err != nil {
		return nil, err
	}
	return &pb.OkResponse{
		Ok: true,
	}, nil
}
