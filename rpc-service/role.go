package rpcservice

import (
	"context"

	pb "github.com/just-arun/micro-session-proto"
	"github.com/just-arun/micro-session/model"
	"github.com/just-arun/micro-session/session"
)

func (r *SessionService) SetRole(ctx context.Context, req *pb.RoleObject) (*pb.OkResponse, error) {
	accesses := []model.Access{}
	for _, v := range req.Access {
		accesses = append(accesses, model.Access{
			ID:   uint(v.Id),
			Name: "",
			Key:  v.Key,
		})
	}
	err := session.RoleAccess().Set(r.GeneralSessionRedisDB, req.Name, &model.Role{
		ID:       uint(req.Id),
		Name:     req.Name,
		Accesses: accesses,
	})
	if err != nil {
		return nil, err
	}
	return &pb.OkResponse{
		Ok: true,
	}, nil
}


func (r *SessionService) GetRole(ctx context.Context, req *pb.GetRoleParam) (*pb.GetRoleReponse, error) {
	accesses := []*pb.AccessObject{}
	data, err := session.RoleAccess().Get(r.GeneralSessionRedisDB, req.Name)
	if err != nil {
		return nil, err
	}
	for _, v := range data.Accesses {
		accesses = append(accesses, &pb.AccessObject{
			Id: uint64(v.ID),
			Key: v.Key,
		})
	}
	return &pb.GetRoleReponse{
		Access: accesses,
	}, nil
}



