package rpcservice

import (
	"io"

	pb "github.com/just-arun/micro-session-proto"
	"github.com/just-arun/micro-session/model"
	"github.com/just-arun/micro-session/session"
)

func (r *SessionService) SetServiceMap(stream pb.SessionService_SetServiceMapServer) error {
	siteMapData := []model.ServiceMap{}
	defer session.SiteMap().Set(r.Ctx.GeneralSessionRedisDB, siteMapData)
	for {
		req, err := stream.Recv()
		if err != io.EOF {
			return stream.SendAndClose(&pb.OkResponse{Ok: false})
		}
		if err != nil {
			return err
		}
		siteMapData = append(siteMapData, model.ServiceMap{
			ID:    uint(req.Id),
			Key:   req.Key,
			Value: req.Value,
			Auth:  req.Auth,
		})
	}
}

func (r *SessionService) GetServiceMap(req *pb.NoPayload, stream pb.SessionService_GetServiceMapServer) error {
	data, err := session.SiteMap().Get(r.Ctx.GeneralSessionRedisDB)
	if err != nil {
		return err
	}
	for _, v := range data {
		err = stream.Send(&pb.ServiceMapData{
			Id:    uint64(v.ID),
			Key:   v.Key,
			Value: v.Value,
			Auth:  v.Auth,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
