/*
backend
accecpt rpc request.Deserialize, log and transmit request to the mircoservice.
*/

package main

import (
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"Clio/component/backend/conf"
	srv "Clio/component/backend/service"
	"Clio/pkg/env"
	pb "Clio/pkg/grpcproto"
	"Clio/pkg/utlis"
)

//backend init, read env var and cmd flags
func initBackEnd() {
	conf.Backend_Host = utils.GetStringEnv(env.BACKEND_HOST, "localhost")
	conf.Backend_Port = utils.GetStringEnv(env.BACKEND_PORT, "8070")

	conf.K8sapihost = utils.GetStringEnv(env.K8S_API_SERVER_HOST, "localhost")
	conf.K8sapiport = utils.GetStringEnv(env.K8S_API_SERVER_PORT, "8080")
	conf.K8sprotocol = utils.GetStringEnv(env.K8S_PROTOCOL, "http")

	conf.UserMhost = utils.GetStringEnv(env.USER_MANAGER_HOST, "127.0.0.1")
	conf.UserMport = utils.GetStringEnv(env.USER_MANAGER_PORT, "8060")

	conf.Debug = flag.Bool("debug", false, "Debug mode, default to false")

	conf.REGSERVERHOST = utils.GetStringEnv(env.REGISTRY_SERVER_HOST, "127.0.0.1")
	conf.REGSERVERPORT = utils.GetStringEnv(env.REGISTRY_SERVER_HOST, "8050")
}

//backend start, listen and start rpc service
func startBackEnd() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", conf.Backend_Host, conf.Backend_Port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterGServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

//create grpc service
func newServer() *srv.GServer {
	s := new(srv.GServer)
	s.InitTimes()
	s.InitSessions()
	return s
}

func main() {
	fmt.Println("backend start")
	initBackEnd()

	flag.Parse()
	if *conf.Debug {
		fmt.Println("debug mode start")
	}
	startBackEnd()

}
