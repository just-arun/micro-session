package rpcservice

import (
	pb "github.com/just-arun/micro-session-proto"
	"github.com/just-arun/micro-session/model"
	"github.com/redis/go-redis/v9"
)

type SessionService struct {
	pb.SessionServiceServer
	GeneralSessionRedisDB *redis.Client
	UserSessionRedisDB    *redis.Client
	Env                   *model.Env
}
