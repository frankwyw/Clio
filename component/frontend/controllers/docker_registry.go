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

type RegController struct {
	beego.Controller
}

func (self *RegController) Reglogin() {
	url := fmt.Sprintf("http://%s:%s", conf.Registry_host, conf.Registry_port)
	uname := ""
	upwd := ""
	is_secure := false
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

	bb := pb.RegLoginReq{Url: url, Username: uname, Password: upwd, Secure: is_secure, Userid: nuserid}

	data, err2 := client.RegLogin(context.Background(), &bb)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(data)

	self.Ctx.WriteString(fmt.Sprintf("%s", data))

}

func (self *RegController) RegListRepo() {
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

	bb := pb.Userid{Userid: nuserid}

	data, err2 := client.RegListRepo(context.Background(), &bb)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(data)

	self.Ctx.WriteString(fmt.Sprintf("%s", data))

}

func (self *RegController) RegGetImageTag() {
	images := ""
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

	bb := pb.RegGetImageTagRep{Images: images, Userid: nuserid}

	data, err2 := client.RegGetImageTag(context.Background(), &bb)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(data)

	self.Ctx.WriteString(fmt.Sprintf("%s", data))
}

func (self *RegController) RegPush() {
	images := ""
	ref := ""
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

	bb := pb.RegImageWithId{Name: images, Reference: ref, Userid: nuserid}

	data, err2 := client.RegPush(context.Background(), &bb)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(data)

	self.Ctx.WriteString(fmt.Sprintf("%s", data))
}

func (self *RegController) RegDelImage() {
	images := ""
	ref := ""
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

	bb := pb.RegImageWithId{Name: images, Reference: ref, Userid: nuserid}

	data, err2 := client.RegDelImage(context.Background(), &bb)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(data)

	self.Ctx.WriteString(fmt.Sprintf("%s", data))
}

func (self *RegController) RegPull() {
	images := ""
	ref := ""
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

	bb := pb.RegImageWithId{Name: images, Reference: ref, Userid: nuserid}

	data, err2 := client.RegDelImage(context.Background(), &bb)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(data)

	self.Ctx.WriteString(fmt.Sprintf("%s", data))
}
