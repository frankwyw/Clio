package controllers

import (
	"context"

	"github.com/astaxie/beego"

	pb "Clio/pkg/grpcproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"Clio/component/frontend/conf"
	"Clio/component/frontend/models"
	"fmt"
)

type TimeController struct {
	beego.Controller
}

//GetTime  get /time?typeid=typeid

func (self *TimeController) Query() {
	typeid, err := self.GetInt32("typeid")
	if err != nil {
		fmt.Println(err)
		return
	}
	userid := self.GetSession("userid")
	if userid == nil {
		return
	}
	var nuserid int32
	var ok bool
	if nuserid, ok = userid.(int32); !ok {
		return
	}
	ok, err = queryTime(typeid, nuserid)
	if err != nil {
		fmt.Println(err)
		return
	}
	self.Ctx.WriteString(fmt.Sprintf("%v", ok))
}

func queryTime(typeid int32, userid int32) (bool, error) {
	serverAddr := fmt.Sprintf("%s:%s", conf.Backend_host, conf.Backend_port)
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
		return false, err
	}
	defer conn.Close()
	client := pb.NewGClient(conn)
	bb := pb.RepTime{ResTypeId: typeid, Userid: userid}
	data, err2 := client.GetResTime(context.Background(), &bb)
	if err2 != nil {
		fmt.Println(err2)
		return false, err2
	}
	fmt.Println(data)
	data_str := fmt.Sprintf("%s", data)
	if models.Times[typeid] != data_str {
		models.Times[typeid] = data_str
		return false, nil
	}
	return true, nil
}
