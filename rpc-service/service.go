package rpcservice

import (
	pb "github.com/just-arun/micro-session-proto"
	"github.com/just-arun/micro-session/model"
)

type SessionService struct {
	pb.SessionServiceServer
	Ctx *model.GlobalCtx
}
