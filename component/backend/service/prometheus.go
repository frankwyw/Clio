package service

import (
	pb "Clio/pkg/grpcproto"
	"net/http"

	"github.com/golang/glog"

	"Clio/pkg/api/prometheus"

	"golang.org/x/net/context"

	"Clio/pkg/utlis"
)

func (s *GServer) PromQuery(ctx context.Context, req *pb.PromeReq) (*pb.Resp, error) {
	userid := req.GetUserid()
	glog.Infof("%d start PromQuery at %s\n", userid, req.GetUrl())
	cfg := prometheus.Config{Address: "127.0.0.1", Transport: prometheus.DefaultTransport}
	cli, err := prometheus.New(cfg)
	if err != nil {
		file, funcn, line := utils.GetPosition()
		glog.Errorf("init PromQuery cfg with error %v, at %s %s %d\n", err, file, funcn, line)
		return nil, err
	}
	client := prometheus.NewApiC(cli)
	url := req.GetUrl()
	req_, _ := http.NewRequest("GET", url, nil)
	resp, body, err := client.Do(ctx, req_)
	if err != nil {
		file, funcn, line := utils.GetPosition()
		glog.Errorf("exec PromQuery with error %v, at %s %s %d, at url %s\n", err, file, funcn, line, req.GetUrl())
		return nil, err
	}

	glog.Infof("%d exec PromQuery success at %s\n", userid, req.GetUrl())
	return &pb.Resp{Httpcode: resp.Status, Resp: string(body)}, nil
}
