package server

import (
	"net"

	"google.golang.org/grpc"

	pb "github.com/just-arun/micro-session-proto"
	"github.com/just-arun/micro-session/boot"
	"github.com/just-arun/micro-session/model"
	"github.com/just-arun/micro-session/pubsub"
	rpcservice "github.com/just-arun/micro-session/rpc-service"
	"github.com/just-arun/micro-session/util"
)

func Run(appEnv, context, port string) {
	con, err := net.Listen("tcp", port)

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	ctxStruct := &rpcservice.SessionService{}
	env := &model.Env{}
	util.GetEnv(".env."+appEnv, context, &env)
	ctx := &model.GlobalCtx{}
	userSessionClient := boot.Redis(env.UserSession.Address, env.UserSession.Password, 0, "User session")
	ctx.UserSessionRedisDB = userSessionClient
	generalSessionClient := boot.Redis(env.GeneralSession.Address, env.GeneralSession.Password, 0, "General session")
	ctx.GeneralSessionRedisDB = generalSessionClient
	ctx.Env = env
	natsCon := boot.NatsConnection(env.Nats.Token)
	ctx.NatsConnection = natsCon
	ctxStruct.Ctx = ctx

	pb.RegisterSessionServiceServer(grpcServer, ctxStruct)
	pubsub.SiteMap(ctxStruct.Ctx).
		SubscribeUpdateSiteMap()

	// go Api(appEnv, port)

	if err := grpcServer.Serve(con); err != nil {
		panic(err)
	}
}
