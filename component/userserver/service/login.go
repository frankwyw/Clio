package service

import (
	"fmt"
	"io"

	restful "github.com/emicklei/go-restful"
)

func LoginSvcConfig() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/login").Produces(restful.MIME_JSON)
	ws.Route(ws.GET("").Param(ws.QueryParameter("username", "username").DataType("string")).To(getLogin))

	return ws
}

//resuful api
//success return json
//error return "error"
func getLogin(req *restful.Request, resp *restful.Response) {
	uname := req.QueryParameter("username")
	if uname == "" {
		fmt.Println("no query param username")
		io.WriteString(resp.ResponseWriter, "error")
		return
	}

	row, err := DB.Query(fmt.Sprintf("SELECT upasswd, uid, uauthority  from users WHERE uname='%s'", uname))
	if err != nil {
		fmt.Println(err)
		io.WriteString(resp.ResponseWriter, "error")
		return
	}

	var pwd string
	var uid int
	var author int
	for row.Next() {
		row.Scan(&pwd, &uid, &author)
	}
	fmt.Println(pwd)
	m := make(map[string]interface{})
	m["password"] = pwd
	m["userid"] = uid
	m["authority"] = author
	resp.WriteEntity(m)
}
