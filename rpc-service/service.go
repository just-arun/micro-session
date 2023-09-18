package rpcservice

import (
	pb "github.com/just-arun/micro-session-proto"
	"github.com/redis/go-redis/v9"
)

type SessionService struct {
	pb.SessionServiceServer
	RedisDB *redis.Client
}
