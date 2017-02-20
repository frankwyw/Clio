package main

import (
	"Clio/pkg/env"
	"Clio/pkg/utlis"
	"Clio/plugin/cicd/conf"
	"Clio/plugin/cicd/service"
	"flag"
	"fmt"
	"net/http"

	restful "github.com/emicklei/go-restful"
)

func initCICDServer() {
	conf.Backend_Host = utils.GetStringEnv(env.BACKEND_HOST, "127.0.0.1")
	conf.Backend_Port = utils.GetStringEnv(env.BACKEND_PORT, "8070")
	conf.Debug = flag.Bool("debug", false, "Debug mode, default to false")

	conf.CICD_HOST = utils.GetStringEnv(conf.CICDHOST, "127.0.0.1")
	conf.CICD_PORT = utils.GetStringEnv(conf.CICDPORT, "8040")

}

func startCICDServer() {
	ws := service.CICDSrvConfig()
	restful.Add(ws)
	server := &http.Server{Addr: fmt.Sprintf("%s:%s", conf.CICD_HOST, conf.CICD_PORT), Handler: restful.DefaultContainer}
	server.ListenAndServe()
}

func main() {
	fmt.Println("cicdserverstart")

	initCICDServer()
	flag.Parse()
	if *conf.Debug {
		fmt.Println("debug mode start")
	}

	startCICDServer()
}
