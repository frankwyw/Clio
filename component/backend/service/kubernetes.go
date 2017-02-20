package service

import (
	pb "Clio/pkg/grpcproto"
	"fmt"

	"github.com/golang/glog"

	"net/http"

	"golang.org/x/net/context"

	"Clio/component/backend/conf"
	"Clio/pkg/utlis"
	"strings"
)

func (s *GServer) K8SRestApi(ctx context.Context, req *pb.K8SReq) (*pb.Resp, error) {
	var resp *http.Response
	var err error
	userid := req.GetUserid()
	glog.Infof("%d start K8SRestApi %s at %s\n", userid, req.GetReqtype(), req.GetUrl())

	if req.Reqtype == "Get" {
		resp, err = getK8s(req.Url)
	} else if req.Reqtype == "Post" {
		resp, err = postK8s(req.Url, req.Bodyparam)
	} else if req.Reqtype == "put" {
		resp, err = putK8s(req.Url, req.Bodyparam)
	} else if req.Reqtype == "delete" {
		resp, err = deleteK8s(req.Url)
	} else {
		err = conf.BackErr{Str: "unknown type"}
		file, funcn, line := utils.GetPosition()
		glog.Errorf("exec K8SRestApi with error %v, at %s %s %d\n", err, file, funcn, line)
	}

	if err != nil {
		return nil, err
	}

	if req.Reqtype != "Get" && (resp.StatusCode/100 == 2) {
		s.times[req.Resint] = fmt.Sprintf("%s", utils.GetNowTime())
	}

	bodystr := utils.Readresp(resp)

	glog.Infof("%d exec K8SRestApi success %s at %s\n", userid, req.GetReqtype(), req.GetUrl())
	return &pb.Resp{Httpcode: resp.Status, Resp: string(bodystr)}, nil

}

func (s *GServer) K8SStreamApi(*pb.K8SReq, pb.G_K8SStreamApiServer) error {
	return nil
}

func getUrl(path string) string {
	url := conf.K8sprotocol + "://" + conf.K8sapihost + ":" + conf.K8sapiport + "/" + path
	return url
}

func getK8s(path string) (*http.Response, error) {
	url := getUrl(path)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		file, funcn, line := utils.GetPosition()
		glog.Errorf("exec getK8s with error %v, at %s %s %d, at url %s\n", err, file, funcn, line, url)
		return nil, err
	}
	return resp, nil
}

func postK8s(path, bodyParam string) (*http.Response, error) {
	url := getUrl(path)
	fmt.Println(url)
	reader := strings.NewReader(bodyParam)
	resp, err := http.Post(url, "application/json", reader)
	if err != nil {
		file, funcn, line := utils.GetPosition()
		glog.Errorf("exec postK8s with error %v, at %s %s %d, url: %s\n", err, file, funcn, line, url)
		return nil, err
	}
	return resp, nil
}

func putK8s(path, bodyParam string) (*http.Response, error) {
	url := getUrl(path)
	fmt.Println(url)
	req, err := http.NewRequest("Put", url, strings.NewReader(bodyParam))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		file, funcn, line := utils.GetPosition()
		glog.Errorf("exec postK8s with error %v, at %s %s %d, url: %s\n", err, file, funcn, line, url)
		return nil, err
	}
	return resp, nil
}

func deleteK8s(path string) (*http.Response, error) {
	url := getUrl(path)
	fmt.Println(url)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		file, funcn, line := utils.GetPosition()
		glog.Errorf("exec getK8s with error %v, at %s %s %d, at url %s\n", err, file, funcn, line, url)
		return nil, err
	}
	return resp, nil
}
