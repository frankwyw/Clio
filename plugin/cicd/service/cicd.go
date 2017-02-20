package service

import (
	"Clio/component/frontend/conf"
	"context"
	"fmt"
	"io"
	"os"
	"strconv"

	pb "Clio/pkg/grpcproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	restful "github.com/emicklei/go-restful"

	"os/exec"

	"io/ioutil"
)

//return ["global"]["workdir"] string
func cdop(m map[string]map[string]interface{}) (string, error) {

	dir := m["global"]["workdir"]

	var ndir string
	var ok bool
	if ndir, ok = dir.(string); !ok {
		return "", nil
	}

	return "cd " + ndir + "\n", nil
}

//cicd control flow
func cicdop(req *restful.Request, resp *restful.Response) {
	userid, err := strconv.Atoi(req.QueryParameter("userid"))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(userid)

	doc := make(map[string]map[string]interface{})
	err = req.ReadEntity(&doc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(doc)

	cdstr, err := cdop(doc)
	if err != nil {
		fmt.Println(err)
		return
	}

	gitstr, err := gitop(doc)
	if err != nil {
		fmt.Println(err)
		return
	}

	complierstr, err := complierop(doc)
	if err != nil {
		fmt.Println(err)
		return
	}

	dockerbulidstr, err := dockerbuildop(doc)
	if err != nil {
		fmt.Println(err)
		return
	}

	nscript := cdstr + "\n" + gitstr + "\n" + complierstr + "\n" + dockerbulidstr + "\n"

	err = op(nscript)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.Remove("Dockerfile")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = pushapi(doc, int32(userid))
	if err != nil {
		fmt.Println(err)
		return
	}

	nscript, err = pushop(doc)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = op(nscript)
	if err != nil {
		fmt.Println(err)
		return
	}

}

//return ["git"] string (be used to script)
func gitop(m map[string]map[string]interface{}) (string, error) {
	addr := m["git"]["addr"]
	script := m["git"]["script"]

	var naddr string
	var ok bool
	if naddr, ok = addr.(string); !ok {
		return "", nil
	}

	gitstr := fmt.Sprintf("git clone %s\n", naddr)

	str, err := scripop(script)
	if err != nil {
		return "", nil
	}

	return gitstr + str, nil
}

//operate string to script and execute
func op(nscript string) error {

	fmt.Println("test", nscript)

	nscript = "#!/bin/bash\n" + nscript

	err := ioutil.WriteFile("nscript", []byte(nscript), 0777)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	scriptcmd := exec.Command("./nscript")

	scriptpip, err := scriptcmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	scriptpip2, err := scriptcmd.StderrPipe()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	str := make([]byte, 200)
	str2 := make([]byte, 200)
	go func() { scriptcmd.Run() }()

	go func() {
		var err error
		for err == nil {
			n, err := scriptpip2.Read(str2)
			if err == nil {
				fmt.Printf("%s", str2[:n])
			} else if err == io.EOF {
				fmt.Printf("%s", str2[:n])
				scriptpip2.Close()
				break
			}
		}
	}()

	for err == nil {
		n, err := scriptpip.Read(str)
		if err == nil {
			fmt.Printf("%s", str[:n])
		} else if err == io.EOF {
			fmt.Printf("%s", str[:n])
			scriptpip.Close()
			break
		}
	}

	err = os.Remove("nscript")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return nil
}

//return ..["script"] string (be used to script)
func scripop(script interface{}) (string, error) {
	var ok bool
	var nscript string
	if nscript, ok = script.(string); !ok {
		return "", nil
	}
	return nscript, nil
}

//return ..["complier"] string (be used to script)
func complierop(m map[string]map[string]interface{}) (string, error) {
	script := m["complier"]["script"]
	return scripop(script)
}

//return ..["docker"] string (be used to script)
func dockerbuildop(m map[string]map[string]interface{}) (string, error) {
	dockerfile := m["docker"]["dockerfile"]
	var ok bool
	var ndockerfile string
	if ndockerfile, ok = dockerfile.(string); !ok {
		return "", nil
	}

	err := ioutil.WriteFile("Dockerfile", []byte(ndockerfile), 0777)
	if err != nil {
		fmt.Println("err")
		return "", nil
	}

	dockerfileparam := m["docker"]["dockerfileparam"]
	var ndockerfileparam string
	if ndockerfileparam, ok = dockerfileparam.(string); !ok {
		return "", nil
	}
	image := m["docker"]["image"]
	var nimage string
	if nimage, ok = image.(string); !ok {
		return "", nil
	}
	ref := m["docker"]["ref"]
	var nref string
	if nref, ok = ref.(string); !ok {
		return "", nil
	}
	imageandref := fmt.Sprintf("%s/%s", nimage, nref)

	dockerstr := fmt.Sprintf("docker build %s %s", ndockerfileparam, imageandref)

	script := m["complier"]["script"]
	str, err := scripop(script)
	return dockerstr + str, err
}

//return ["push"]["script"] string (be used to script)
func pushop(m map[string]map[string]interface{}) (string, error) {
	script := m["push"]["script"]
	return scripop(script)
}

//push operate after docker build
func pushapi(m map[string]map[string]interface{}, userid int32) (*io.ReadCloser, error) {

	regisrtyaddr := m["docker"]["regisrtyaddr"]
	var ok bool
	var nregisrtyaddr string
	if nregisrtyaddr, ok = regisrtyaddr.(string); !ok {
		return nil, nil
	}

	username := m["docker"]["username"]
	var nusername string
	if nusername, ok = username.(string); !ok {
		return nil, nil
	}

	password := m["docker"]["password"]
	var npassword string
	if npassword, ok = password.(string); !ok {
		return nil, nil
	}

	secure := m["docker"]["secure"]
	var nsecure bool
	if nsecure, ok = secure.(bool); !ok {
		return nil, nil
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

	image := m["docker"]["image"]
	ref := m["docker"]["ref"]

	aa := pb.RegLoginReq{Url: nregisrtyaddr, Username: nusername, Password: npassword, Secure: nsecure, Userid: userid}
	data0, err1 := client.RegLogin(context.Background(), &aa)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(data0)

	bb := pb.RegImageWithId{Name: fmt.Sprintf("%v", image), Reference: fmt.Sprintf("%s", ref), Userid: userid}

	data, err2 := client.RegPush(context.Background(), &bb)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(data)

	return nil, nil
}

func CICDSrvConfig() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/cicd")
	ws.Route(ws.POST("").Param(ws.QueryParameter("userid", "identify user").DataType("int")).To(cicdop))
	return ws
}
