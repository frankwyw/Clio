package controllers

import (
	"context"

	"github.com/astaxie/beego"

	"Clio/component/frontend/conf"
	pb "Clio/pkg/grpcproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"fmt"
)

type UserMController struct {
	beego.Controller
}

func (self *UserMController) GetAllUsers() {
	var url string = "users?authority="
	userid := self.GetSession("userid")
	if userid == nil {
		return
	}
	var nuserid int32
	var ok bool
	if nuserid, ok = userid.(int32); !ok {
		return
	}
	UserMreq("", url, "Get", nuserid)
}

func (self *UserMController) GetUser() {
	var url string = "users/1?authority="
	userid := self.GetSession("userid")
	if userid == nil {
		return
	}
	var nuserid int32
	var ok bool
	if nuserid, ok = userid.(int32); !ok {
		return
	}
	UserMreq("", url, "Get", nuserid)
}

func (self *UserMController) Post() {
	url := "users"
	userid := self.GetSession("userid")
	if userid == nil {
		return
	}
	var nuserid int32
	var ok bool
	if nuserid, ok = userid.(int32); !ok {
		return
	}
	UserMreq(fmt.Sprintf("%s", self.Ctx.Input.RequestBody), url, "Post", nuserid)
}

func (self *UserMController) Put() {
	var url string = "/users/3?authority="
	userid := self.GetSession("userid")
	if userid == nil {
		return
	}
	var nuserid int32
	var ok bool
	if nuserid, ok = userid.(int32); !ok {
		return
	}
	UserMreq(fmt.Sprintf("%s", self.Ctx.Input.RequestBody), url, "Put", nuserid)
}

func (self *UserMController) Delete() {
	var url string = " /users/1"
	userid := self.GetSession("userid")
	if userid == nil {
		return
	}
	var nuserid int32
	var ok bool
	if nuserid, ok = userid.(int32); !ok {
		return
	}
	UserMreq("", url, "delete", nuserid)
}

func UserMreq(bodyparam, url, reqm string, userid int32) (*pb.Resp, error) {
	serverAddr := fmt.Sprintf("%s:%s", conf.Backend_host, conf.Backend_port)
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
		return nil, err
	}
	defer conn.Close()
	client := pb.NewGClient(conn)
	bb := pb.UseMrRep{Reqmethod: reqm, Requrl: url, Userid: userid, Bodyparam: bodyparam}
	data, err := client.UserMApi(context.Background(), &bb)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}
