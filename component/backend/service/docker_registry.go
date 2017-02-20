package service

import (
	pb "Clio/pkg/grpcproto"
	"fmt"
	"net/http"

	"golang.org/x/net/context"

	"Clio/component/backend/conf"

	"github.com/golang/glog"

	"Clio/pkg/utlis"
)

//registry login rpc
//success return resp
//error return error
func (s *GServer) RegLogin(ctx context.Context, req *pb.RegLoginReq) (*pb.Resp, error) {
	url := req.GetUrl()
	username := req.GetUsername()
	passwd := req.GetPassword()
	userid := req.GetUserid()
	secure := req.GetSecure()
	glog.Infof("%d start RegLogin\n", userid)

	resp, err := http.Get(fmt.Sprintf("%s:%s/registry/login?url=%s&username=%s&password=%s&secure=%v&userid=%d",
		conf.REGSERVERHOST, conf.REGSERVERPORT, url, username, passwd, secure, userid))

	if err != nil {
		glog.Errorf("%d exec RegLogin with error %v\n", userid, err)
	}

	bodystr := utils.Readresp(resp)

	var respstr string
	if resp.StatusCode == 200 {
		respstr = string(bodystr)
	} else {
		respstr = string(bodystr)
	}

	glog.Infof("%d exec RegLogin sucess\n", userid)

	return &pb.Resp{Httpcode: resp.Status, Resp: respstr}, nil
}

//regsitry list all repo after login
//success return resp
//error return error
func (s *GServer) RegListRepo(ctx context.Context, req *pb.Userid) (*pb.Resp, error) {
	userid := req.GetUserid()
	glog.Infof("%d start RegListRepo\n", userid)

	resp, err := http.Get(fmt.Sprintf("%s:%s/registry/repo?userid=%d",
		conf.REGSERVERHOST, conf.REGSERVERPORT, userid))

	if err != nil {
		glog.Errorf("%d exec RegListRepo with error %v\n", userid, err)
	}

	bodystr := utils.Readresp(resp)

	var respstr string
	if resp.StatusCode == 200 {
		respstr = string(bodystr)
	} else {
		respstr = string(bodystr)
	}

	glog.Infof("%d exec RegListRepo sucess\n", userid)

	return &pb.Resp{Httpcode: resp.Status, Resp: respstr}, nil
}

//regsitry get image tag after login
//success return resp
//error return error
func (s *GServer) RegGetImageTag(ctx context.Context, req *pb.RegGetImageTagRep) (*pb.Resp, error) {
	userid := req.GetUserid()
	images := req.GetImages()
	glog.Infof("%d start RegGetImageTag %s\n", userid, images)
	resp, err := http.Get(fmt.Sprintf("%s:%s/registry/tag?userid=%d&image=%s",
		conf.REGSERVERHOST, conf.REGSERVERPORT, userid, images))

	if err != nil {
		glog.Errorf("%d exec RegGetImageTag with error %v\n", userid, err)
	}

	bodystr := utils.Readresp(resp)

	var respstr string
	if resp.StatusCode == 200 {
		respstr = string(bodystr)
	} else {
		respstr = string(bodystr)
	}

	glog.Infof("%d exec RegGetImageTag sucess\n", userid)

	return &pb.Resp{Httpcode: resp.Status, Resp: respstr}, nil
}

//push the image where the peer has to registry after login
//success return resp
//error return error
func (s *GServer) RegPush(ctx context.Context, req *pb.RegImageWithId) (*pb.Resp, error) {
	userid := req.GetUserid()
	image := req.GetName()
	ref := req.GetReference()
	glog.Infof("%d start RegPush\n", userid)

	resp, err := http.Get(fmt.Sprintf("%s:%s/registry/push?userid=%d&image=%s&ref=%s",
		conf.REGSERVERHOST, conf.REGSERVERPORT, userid, image, ref))

	if err != nil {
		glog.Errorf("%d exec RegPush with error %v\n", userid, err)
	}

	bodystr := utils.Readresp(resp)

	var respstr string
	if resp.StatusCode == 200 {
		respstr = string(bodystr)
	} else {
		respstr = string(bodystr)
	}

	glog.Infof("%d exec RegPush sucess\n", userid)

	return &pb.Resp{Httpcode: resp.Status, Resp: respstr}, nil
}

//delete the image in the registry after login
//success return resp
//error return error
func (s *GServer) RegDelImage(ctx context.Context, req *pb.RegImageWithId) (*pb.Resp, error) {
	userid := req.GetUserid()
	image := req.GetName()
	ref := req.GetReference()
	glog.Infof("%d start RegDelImage\n", userid)

	resp, err := http.Get(fmt.Sprintf("%s:%s/registry/del?userid=%d&image=%s&ref=%s",
		conf.REGSERVERHOST, conf.REGSERVERPORT, userid, image, ref))

	if err != nil {
		glog.Errorf("%d exec RegDelImage with error %v\n", userid, err)
	}

	bodystr := utils.Readresp(resp)

	var respstr string
	if resp.StatusCode == 200 {
		respstr = string(bodystr)
	} else {
		respstr = string(bodystr)
	}

	glog.Infof("%d exec RegDelImage sucess\n", userid)

	return &pb.Resp{Httpcode: resp.Status, Resp: respstr}, nil
}

//pull the image from registry to the peer after login
//success return resp
//error return error
func (s *GServer) RegPull(ctx context.Context, req *pb.RegImageWithId) (*pb.Resp, error) {
	userid := req.GetUserid()
	image := req.GetName()
	ref := req.GetReference()
	glog.Infof("%d start RegPull\n", userid)

	resp, err := http.Get(fmt.Sprintf("%s:%s/registry/pull?userid=%d&image=%s&ref=%s",
		conf.REGSERVERHOST, conf.REGSERVERPORT, userid, image, ref))

	if err != nil {
		glog.Errorf("%d exec RegPull with error %v\n", userid, err)
	}

	bodystr := utils.Readresp(resp)

	var respstr string
	if resp.StatusCode == 200 {
		respstr = string(bodystr)
	} else {
		respstr = string(bodystr)
	}

	glog.Infof("%d exec RegPull sucess\n", userid)

	return &pb.Resp{Httpcode: resp.Status, Resp: respstr}, nil
}
