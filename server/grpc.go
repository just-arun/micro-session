package server

import (
	"net"

	"google.golang.org/grpc"

	pb "github.com/just-arun/micro-session-proto"
	rpcService "github.com/just-arun/micro-session/rpc-service"
)

func Run(port string) {
	con, err := net.Listen("tcp", port)

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterSessionServiceServer(grpcServer, &rpcService.SessionService{})

	if err := grpcServer.Serve(con); err != nil {
		panic(err)
	}
}
