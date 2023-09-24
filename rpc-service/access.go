package rpcservice

import (
	"context"

	pb "github.com/just-arun/micro-session-proto"
	"github.com/just-arun/micro-session/session"
	"github.com/just-arun/micro-session/util"
)

func (r *SessionService) HaveAccess(ctx context.Context, req *pb.HaveAccessParam) (*pb.HaveAccessResponse, error) {
	access, err := session.RoleAccess().GetAccessesForRoles(r.GeneralSessionRedisDB, req.Roles)
	if err != nil {
		return nil, err
	}
	hasRole := util.Array().Includes(access, func(item string, _ int) bool {
		return item == req.AccessSlug
	})
	return &pb.HaveAccessResponse{
		Access: hasRole,
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
