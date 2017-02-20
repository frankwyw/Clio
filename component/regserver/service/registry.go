package service

import (
	"fmt"
	"io"
	"os/exec"
	"strconv"

	restful "github.com/emicklei/go-restful"
	"github.com/golang/glog"

	"github.com/heroku/docker-registry-client/registry"
)

//init the registrt client sessions
func init() {
	sessions = make(map[int]*registry.Registry)
}

var (
	sessions map[int]*registry.Registry
)

//restful api
//success return "login success"
//error return "error"
func loginReg(req *restful.Request, resp *restful.Response) {
	url := req.QueryParameter("url")
	username := req.QueryParameter("username")
	password := req.QueryParameter("password")
	//	secure := req.QueryParameter("secure")
	userid, err := strconv.Atoi(req.QueryParameter("userid"))

	if err != nil {
		fmt.Println(err)
		io.WriteString(resp.ResponseWriter, "error")
		return
	}

	glog.Infof("%d start loginReg\n", userid)

	if url == "" {
		fmt.Println("error url")
		io.WriteString(resp.ResponseWriter, "error")
		return
	}

	cli, err := registry.New(url, username, password)
	if err != nil {
		glog.Errorf("%d exec RegLogin with error %v\n", userid, err)
		io.WriteString(resp.ResponseWriter, "error")
		return
	}

	sessions[userid] = cli

	glog.Infof("%d exec loginReg success\n", userid)

	io.WriteString(resp.ResponseWriter, "login success")
}

//restful api
//success return repos
//error return "error"
func listRepo(req *restful.Request, resp *restful.Response) {
	userid, err := strconv.Atoi(req.QueryParameter("userid"))
	if err != nil {
		fmt.Println(err)
		io.WriteString(resp.ResponseWriter, "error")
		return
	}

	glog.Infof("%d start listRepo\n", userid)

	cli := sessions[userid]
	if cli == nil {
		fmt.Println("registry no login")
		io.WriteString(resp.ResponseWriter, "error")
		return
	}

	repositories, err := cli.Repositories()

	if err != nil {
		glog.Errorf("%d exec RegListRepo with error %v\n", userid, err)
		io.WriteString(resp.ResponseWriter, "error")
		return
	}

	glog.Infof("%d exec listRepo success\n", userid)

	io.WriteString(resp.ResponseWriter, fmt.Sprintf("%v", repositories))

}

//restful api
//success return tag
//error return "error"
func getTag(req *restful.Request, resp *restful.Response) {
	userid, err := strconv.Atoi(req.QueryParameter("userid"))
	if err != nil {
		fmt.Println(err)
		io.WriteString(resp.ResponseWriter, "error")
		return
	}

	glog.Infof("%d start getTag\n", userid)

	cli := sessions[userid]
	if cli == nil {
		fmt.Println("registry no login")
		io.WriteString(resp.ResponseWriter, "error")
		return
	}

	iname := req.QueryParameter("image")

	tags, err := cli.Tags(iname)

	if err != nil {
		glog.Errorf("%d exec getTag with error %v\n", userid, err)
		io.WriteString(resp.ResponseWriter, "error")
		return
	}

	fmt.Println(tags)

	glog.Infof("%d exec getTag success\n", userid)

	io.WriteString(resp.ResponseWriter, fmt.Sprintf("%v", tags))

}

//restful api
//success return "push success"
//error return "error"
func pushImage(req *restful.Request, resp *restful.Response) {
	userid, err := strconv.Atoi(req.QueryParameter("userid"))
	if err != nil {
		fmt.Println(err)
		io.WriteString(resp.ResponseWriter, "error")
		return
	}

	glog.Infof("%d start pushImage\n", userid)

	cli := sessions[userid]
	if cli == nil {
		fmt.Println("registry no login")
		io.WriteString(resp.ResponseWriter, "error")
		return
	}

	image := fmt.Sprintf("%s:%s", req.QueryParameter("image"), req.QueryParameter("ref"))

	cmd := exec.Command("docker", "push", image)
	out, err := cmd.StdoutPipe()

	if err != nil {
		glog.Errorf("%d exec pushImage with error %v\n", userid, err)
		io.WriteString(resp.ResponseWriter, "error")
		return
	}

	str := make([]byte, 200)
	go func() { cmd.Run() }()

	for err == nil {
		n, err := out.Read(str)
		if err == nil {
			fmt.Printf("%s", str[:n])
		} else if err == io.EOF {
			fmt.Printf("%s", str[:n])
			out.Close()
			break
		}
	}

	glog.Infof("%d exec pushImage success\n", userid)

	//should use stream
	io.WriteString(resp.ResponseWriter, "push success")

}

//restful api
//success return "pull success"
//error return "error"
func pullImage(req *restful.Request, resp *restful.Response) {
	userid, err := strconv.Atoi(req.QueryParameter("userid"))
	if err != nil {
		fmt.Println(err)
		io.WriteString(resp.ResponseWriter, "error")
		return
	}

	glog.Infof("%d start pullImage\n", userid)

	cli := sessions[userid]
	if cli == nil {
		fmt.Println("registry no login")
		io.WriteString(resp.ResponseWriter, "error")
		return
	}

	image := fmt.Sprintf("%s:%s", req.QueryParameter("image"), req.QueryParameter("ref"))

	cmd := exec.Command("docker", "pull", image)
	out, err := cmd.StdoutPipe()

	if err != nil {
		glog.Errorf("%d exec pullImage with error %v\n", userid, err)
		io.WriteString(resp.ResponseWriter, "error")
		return
	}

	str := make([]byte, 200)
	go func() { cmd.Run() }()

	for err == nil {
		n, err := out.Read(str)
		if err == nil {
			fmt.Printf("%s", str[:n])
		} else if err == io.EOF {
			fmt.Printf("%s", str[:n])
			out.Close()
			break
		}
	}

	glog.Infof("%d exec pullImage success\n", userid)

	//should use stream
	io.WriteString(resp.ResponseWriter, "pull success")

}

//restful api
//success return "del success"
//error return "error"
func delImage(req *restful.Request, resp *restful.Response) {
	userid, err := strconv.Atoi(req.QueryParameter("userid"))
	if err != nil {
		fmt.Println(err)
		io.WriteString(resp.ResponseWriter, "error")
		return
	}

	glog.Infof("%d start delImage\n", userid)

	cli := sessions[userid]
	if cli == nil {
		fmt.Println("registry no login")
		io.WriteString(resp.ResponseWriter, "error")
		return
	}

	digest, err := cli.ManifestDigest(req.QueryParameter("image"), req.QueryParameter("ref"))
	if err != nil {
		glog.Errorf("%d exec RegDelImage with error %v", userid, err)
		io.WriteString(resp.ResponseWriter, "error")
		return
	}
	err = cli.DeleteManifest(req.PathParameter("image"), digest)
	if err != nil {
		glog.Errorf("%d exec RegDelImage with error %v", userid, err)
		io.WriteString(resp.ResponseWriter, "error")
		return
	}
	glog.Infof("%d exec delImage success\n", userid)

	io.WriteString(resp.ResponseWriter, "del success")

}

func RegSrvConfig() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/registry")
	ws.Route(ws.GET("/login").Param(ws.QueryParameter("url", "registry url").DataType("string")).
		Param(ws.QueryParameter("username", "registry username").DataType("string")).
		Param(ws.QueryParameter("password", "registry password").DataType("string")).
		Param(ws.QueryParameter("secure", "registry is secure").DataType("bool")).
		Param(ws.QueryParameter("userid", "identify user").DataType("int")).
		To(loginReg))
	ws.Route(ws.GET("/repo").
		Param(ws.QueryParameter("userid", "identify user").DataType("int")).To(listRepo))
	ws.Route(ws.GET("/tag").
		Param(ws.QueryParameter("image", "image name").DataType("string")).
		Param(ws.QueryParameter("userid", "identify user").DataType("int")).To(getTag))
	ws.Route(ws.GET("/push").
		Param(ws.QueryParameter("image", "image name").DataType("string")).
		Param(ws.QueryParameter("ref", "image ref").DataType("string")).
		Param(ws.QueryParameter("userid", "identify user").DataType("int")).To(pushImage))
	ws.Route(ws.GET("/pull").
		Param(ws.QueryParameter("image", "image name").DataType("string")).
		Param(ws.QueryParameter("ref", "image ref").DataType("string")).
		Param(ws.QueryParameter("userid", "identify user").DataType("int")).To(pullImage))
	ws.Route(ws.GET("/del").
		Param(ws.QueryParameter("image", "image name").DataType("string")).
		Param(ws.QueryParameter("ref", "image ref").DataType("string")).
		Param(ws.QueryParameter("userid", "identify user").DataType("int")).To(delImage))

	return ws
}
