package main

import (
	"Clio/component/regserver/conf"
	"Clio/component/regserver/service"
	"Clio/pkg/env"
	"Clio/pkg/utlis"
	"flag"
	"fmt"
	"net/http"

	restful "github.com/emicklei/go-restful"
)

func initRegServer() {
	conf.REGSERVERHOST = utils.GetStringEnv(env.REGISTRY_SERVER_HOST, "127.0.0.1")
	conf.REGSERVERPORT = utils.GetStringEnv(env.REGISTRY_SERVER_HOST, "8050")
	conf.Debug = flag.Bool("debug", false, "Debug mode, default to false")

	conf.REGHOST = utils.GetStringEnv(env.REGISTRY_HOST, "127.0.0.1")
	conf.REGPORT = utils.GetStringEnv(env.REGISTRY_PORT, "5000")

}

func startRegServer() {
	ws := service.RegSrvConfig()
	restful.Add(ws)
	server := &http.Server{Addr: fmt.Sprintf("%s:%s", conf.REGSERVERHOST, conf.REGSERVERPORT), Handler: restful.DefaultContainer}
	server.ListenAndServe()
}

func main() {
	fmt.Println("regserverstart")

	initRegServer()
	flag.Parse()
	if *conf.Debug {
		fmt.Println("debug mode start")
	}

	startRegServer()
}
