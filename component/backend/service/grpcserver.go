package service

import (
	"Clio/component/backend/conf"
	pb "Clio/pkg/grpcproto"
	"Clio/pkg/utlis"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/net/context"

	"github.com/golang/glog"
	"github.com/heroku/docker-registry-client/registry"
)

const (
	TIMES_LEN           = 18
	K8S_API_SERVER_HOST = "K8S_API_SERVER_HOST"
	K8S_API_SERVER_PORT = "K8S_API_SERVER_PORT"

	CookiesInit = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	CookiesLen  = 25
)

type userSession struct {
	userId    int
	authority int
	hcli      *HClient
}

type HClient struct {
	reg *registry.Registry
}

type GServer struct {
	/*
	   resource update_time
	   0:  Pod
	   1:  PodTemplate
	   2:  ReplicationController
	   3:  Service
	   4:  EndPoint
	   5:  Node
	   6:  Binding
	   7:  Event
	   8:  LimitRange
	   9:  ResourceQuota
	   10: NameSpace
	   11: Secret
	   12: ServiceAccount
	   13: PersistentVolume
	   14: PersistentVolumeClaim
	   15: DeleteOptions
	   16: ComponentStatus
	   17: ConfigMap
	*/
	times [TIMES_LEN]string
	//	sessions map[string]*userSession
	authormap map[int32]int
	loginIds  []int
}

type loginre struct {
	Password  string
	Userid    int
	Authority int
}

func (s *GServer) InitSessions() {
	//	s.sessions = make(map[string]*userSession)
	//  for test
	//	s.sessions["test"] = &userSession{userId: 4, authority: 0}
	s.authormap = make(map[int32]int)
}

func (s *GServer) InitTimes() {
	time_ := fmt.Sprintf("%s", utils.GetNowTime())
	for i := 0; i < TIMES_LEN; i++ {
		s.times[i] = time_
	}
}

func (s *GServer) GetResTime(ctx context.Context, req *pb.RepTime) (*pb.RespTime, error) {
	pos := req.ResTypeId
	userid := req.GetUserid()
	glog.Infof("%d start GetResTime\n", userid)

	defer func() {
		if err := recover(); err != nil {
			file, funcn, line := utils.GetPosition()
			glog.Errorf("%d exec GetResTime with error %v, at %s %s %d\n", userid, err, file, funcn, line)
		}
	}()
	time_ := s.times[pos]
	glog.Infof("%d exec GetResTime success\n", userid)
	return &pb.RespTime{Time: time_}, nil
}

func (s *GServer) Login(ctx context.Context, req *pb.UserInfo) (*pb.Resp, error) {
	uname := req.GetUser()
	upasswd := req.GetPasswd()
	path := "login" + "?username=" + uname
	glog.Infof("%d start login", uname)

	url := fmt.Sprintf("http://%s:%s/%s", conf.UserMhost, conf.UserMport, path)
	var resp *http.Response
	var err error
	fmt.Println(url)
	resp, err = http.Get(url)
	if err != nil {
		return nil, err
	}

	bodystr := utils.Readresp(resp)
	//get userid, authority, upasswd_
	re := new(loginre)

	err = json.Unmarshal(bodystr, re)
	if err != nil {
		fmt.Println(err)
	}

	var respstr string

	if re.Password == upasswd {
		respstr = fmt.Sprintf("login success\n%d\n", re.Userid)
		//		s.assignsession(re.Userid, re.Authority)
	} else {
		respstr = "login error\n"
	}

	glog.Infof("%d exec login success\n", userid)

	return &pb.Resp{Httpcode: resp.Status, Resp: respstr}, nil
}

/*
func (s *GServer) assignsession(userid, author int) {
	cookie := getRandStr(CookiesLen)
	done := make(chan struct{})
	if len(s.loginIds) == 0 {
		close(done)
	}
	for _, i := range s.loginIds {
		if i == int(userid) {
			go func() {
				for k, j := range s.sessions {
					if j.userId == int(userid) {
						glog.Infof("%d exec RegisterClient with repeated cookies\n", userid)
						delete(s.sessions, k)
					}
				}
				close(done)
			}()
		} else {
			close(done)
		}
	}
	<-done
	s.sessions[cookie] = &userSession{userId: int(userid), authority: author}
}


func getRandStr(n int) string {
	tmp := make([]byte, n)
	length := len(CookiesInit)
	for i := range tmp {
		tmp[i] = CookiesInit[rand.Intn(length)]
	}
	return string(tmp)
}
*/
