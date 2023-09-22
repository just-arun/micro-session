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
