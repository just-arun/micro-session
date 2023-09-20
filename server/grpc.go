package server

import (
	"net"

	"google.golang.org/grpc"

	pb "github.com/just-arun/micro-session-proto"
	"github.com/just-arun/micro-session/boot"
	"github.com/just-arun/micro-session/model"
	rpcservice "github.com/just-arun/micro-session/rpc-service"
	"github.com/just-arun/micro-session/util"
)

func Run(appEnv, port string) {
	con, err := net.Listen("tcp", port)

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	ctxStruct := &rpcservice.SessionService{}
	env := &model.Env{}
	util.GetEnv(".env."+appEnv, ".", &env)
	userSessionClient := boot.Redis(env.UserSession.Address, env.UserSession.Password, 0, "User session")
	ctxStruct.UserSessionRedisDB = userSessionClient
	generalSessionClient := boot.Redis(env.GeneralSession.Address, env.GeneralSession.Password, 0, "General session")
	ctxStruct.GeneralSessionRedisDB = generalSessionClient
	ctxStruct.Env = env

	pb.RegisterSessionServiceServer(grpcServer, ctxStruct)

	if err := grpcServer.Serve(con); err != nil {
		panic(err)
	}
}
