package controllers

import (
	"context"

	"github.com/astaxie/beego"

	"Clio/component/frontend/conf"
	"Clio/pkg/env"
	pb "Clio/pkg/grpcproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"fmt"
)

type K8sController struct {
	beego.Controller
}

func (self *K8sController) Get() {
	var url string = "api/v1/namespaces/default/pods"
	object := env.NAMESPACE
	userid := self.GetSession("userid")
	if userid == nil {
		return
	}
	var nuserid int32
	var ok bool
	if nuserid, ok = userid.(int32); !ok {
		return
	}
	k8sreqhelp("", url, "Get", nuserid, object)
}

func (self *K8sController) Post() {
	url := "api/v1/namespaces/default/pods"
	object := env.NAMESPACE
	userid := self.GetSession("userid")
	if userid == nil {
		return
	}
	var nuserid int32
	var ok bool
	if nuserid, ok = userid.(int32); !ok {
		return
	}
	k8sreqhelp(fmt.Sprintf("%s", self.Ctx.Input.RequestBody), url, "Post", nuserid, object)
}

func (self *K8sController) Put() {
	var url string = "api/v1/namespaces/default/pods"
	object := env.NAMESPACE
	userid := self.GetSession("userid")
	if userid == nil {
		return
	}
	var nuserid int32
	var ok bool
	if nuserid, ok = userid.(int32); !ok {
		return
	}
	k8sreqhelp(fmt.Sprintf("%s", self.Ctx.Input.RequestBody), url, "Put", nuserid, object)
}

func (self *K8sController) Delete() {
	var url string = "api/v1/namespaces/default/pods"
	object := env.NAMESPACE
	userid := self.GetSession("userid")
	if userid == nil {
		return
	}
	var nuserid int32
	var ok bool
	if nuserid, ok = userid.(int32); !ok {
		return
	}
	k8sreqhelp("", url, "delete", nuserid, object)
}

func k8sreqhelp(bodyparam, url, reqm string, userid int32, object int) (*pb.Resp, error) {
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

	bb := pb.K8SReq{Resint: int32(object), Reqtype: reqm, Url: url, Bodyparam: bodyparam, Userid: userid}
	data, err := client.K8SRestApi(context.Background(), &bb)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}
