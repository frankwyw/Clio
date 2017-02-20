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

type PromController struct {
	beego.Controller
}

func (self *PromController) PromQuery() {
	url := ""
	is_range := false
	userid := self.GetSession("userid")
	if userid == nil {
		return
	}
	var nuserid int32
	var ok bool
	if nuserid, ok = userid.(int32); !ok {
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

	bb := pb.PromeReq{Is_Range: is_range, Url: url, Userid: nuserid}
	data, err2 := client.PromQuery(context.Background(), &bb)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(data)
}
