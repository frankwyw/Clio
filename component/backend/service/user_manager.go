package service

import (
	"Clio/component/backend/conf"
	pb "Clio/pkg/grpcproto"
	"Clio/pkg/utlis"
	"fmt"
	"net/http"

	"github.com/golang/glog"

	"golang.org/x/net/context"

	"strconv"
	"strings"
)

func (s *GServer) UserMApi(ctx context.Context, req *pb.UseMrRep) (*pb.Resp, error) {
	userid := req.GetUserid()
	reqmethod := req.GetReqmethod()
	bodyparam := req.GetBodyparam()
	path := req.GetRequrl()
	glog.Infof("%d start UserMApi , %s at %s\n", userid, reqmethod, path)
	authority := s.authormap[userid]

	if strings.Contains(path, "authority=") {
		path = fmt.Sprintf("%s%d", path, authority)
	}

	var resp *http.Response
	var err error

	if reqmethod == "Get" || reqmethod == "Delete" {
		resp, err = getordelUserM(reqmethod, path, authority)
	} else if reqmethod == "Post" || reqmethod == "Put" {

		authority_, err := getCreatedAuthor(bodyparam)
		if err != nil {
			return nil, err
		}

		if authority > authority_ {
			return nil, conf.BackErr{Str: "low authority"}
		}

		resp, err = postorputUserM(reqmethod, path, bodyparam, authority)
	}

	bodystr := utils.Readresp(resp)

	glog.Infof("%d exec UserMApi success, %s at %s\n", userid, reqmethod, path)
	return &pb.Resp{Httpcode: resp.Status, Resp: string(bodystr)}, nil
}

func getCreatedAuthor(bodyparam string) (int, error) {
	n := strings.Index(bodyparam, "\"authority\"")
	op := strings.Index(bodyparam[n:], ":") + 1 + n
	ed := strings.Index(bodyparam[op:], ",") + op
	str := bodyparam[op:ed]
	str = strings.Trim(str, " ")
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func getordelUserM(reqm, path string, authority int) (*http.Response, error) {
	url := fmt.Sprintf("http://%s:%s/%s", conf.UserMhost, conf.UserMport, path)
	var resp *http.Response
	var err error
	if reqm == "Get" {
		resp, err = http.Get(url)
		if err != nil {
			return nil, err
		}
	} else {
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			return nil, err
		}
		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil

}

func postorputUserM(reqm, path, bodypapram string, authority int) (*http.Response, error) {
	url := fmt.Sprintf("http://%s:%s/%s?authority=%d", conf.UserMhost, conf.UserMport, path, authority)
	var resp *http.Response
	var err error
	if reqm == "Post" {
		resp, err = http.Post(url, "application/json", strings.NewReader(bodypapram))
		if err != nil {
			return nil, err
		}
	} else {
		req, err := http.NewRequest("Put", url, strings.NewReader(bodypapram))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}
