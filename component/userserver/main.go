package main

import (
	"Clio/component/userserver/conf"
	"Clio/component/userserver/service"
	"Clio/pkg/env"
	"Clio/pkg/utlis"
	"database/sql"
	"flag"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	restful "github.com/emicklei/go-restful"
)

//userserver init
//read from env var
func initUserMServer() {
	conf.UserMhost = utils.GetStringEnv(env.USER_MANAGER_HOST, "127.0.0.1")
	conf.UserMport = utils.GetStringEnv(env.USER_MANAGER_PORT, "8060")
	conf.Debug = flag.Bool("debug", false, "Debug mode, default to false")

	conf.DBuser = flag.String("dbuser", "root", "mysql username, default root")
	conf.DBpasswd = flag.String("dbpasswd", "", "mysql password, default empty")
	conf.DBname = flag.String("dbname", "test", "mysql databasename, default test")

	conf.UserDBhost = utils.GetStringEnv(env.USER_DB_HOST, "127.0.0.1")
	conf.UserDBport = utils.GetStringEnv(env.USER_DB_PORT, "3306")
}

//start server
func startUserMServer() {
	var err error
	service.DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s", *conf.DBuser, *conf.DBpasswd, conf.UserDBhost, conf.UserDBport, *conf.DBname))
	if err != nil {
		fmt.Println(err)
		return
	}
	ws := service.UserSrvConfig()
	ws2 := service.LoginSvcConfig()
	restful.Add(ws)
	restful.Add(ws2)
	server := &http.Server{Addr: fmt.Sprintf("%s:%s", conf.UserMhost, conf.UserMport), Handler: restful.DefaultContainer}
	server.ListenAndServe()
}

func main() {
	fmt.Println("userserverstart")

	initUserMServer()
	flag.Parse()
	if *conf.Debug {
		fmt.Println("debug mode start")
	}

	startUserMServer()
}
