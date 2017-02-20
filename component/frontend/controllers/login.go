package controllers

import (
	"context"
	"strconv"
	"strings"

	"github.com/astaxie/beego"

	"Clio/component/frontend/conf"
	pb "Clio/pkg/grpcproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"fmt"
)

type LoginController struct {
	beego.Controller
}

func (self *LoginController) Login() {
	user := ""
	pwd := ""

	if user == "" || pwd == "" {
		return
	}

	serverAddr := fmt.Sprintf("%s:%s", conf.Backend_host, conf.Backend_port)
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewGClient(conn)

	bb := pb.UserInfo{User: user, Passwd: pwd}
	data, err2 := client.Login(context.Background(), &bb)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(data)
	strs := strings.Split(data.GetResp(), "\n")
	if strs[0] == "login success" {
		id, err := strconv.Atoi(strs[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		self.SetSession("userid", int(id))
	}
}
