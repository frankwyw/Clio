package service

import (
	"fmt"
	"io"
	"strconv"

	restful "github.com/emicklei/go-restful"

	"database/sql"
)

var (
	DB *sql.DB
)

type user struct {
	uid         string
	uname       string
	uauthority  int
	ucreatetime string
	ucreator    int
}

//get id and authority from Request
func getIdAndAuthor(req *restful.Request) (int, int, error) {
	userId, err := strconv.Atoi(req.PathParameter("user-id"))
	if err != nil {
		fmt.Println(err)
		return -1, -1, err
	}
	authority, err := strconv.Atoi(req.QueryParameter("authority"))
	if err != nil {
		fmt.Println(err)
		return -1, -1, err
	}
	fmt.Println(userId, authority)

	return userId, authority, nil
}

//strcut user to map (be used to json)
func usertomap(u *user, m map[string]interface{}) map[string]interface{} {
	m["userid"] = u.uid
	m["username"] = u.uname
	m["authority"] = u.uauthority
	m["createtime"] = u.ucreatetime
	m["creator"] = u.ucreator

	return m
}

//resuful api, get the user info that has low authority
//return json
//error return "error"
func getUser(req *restful.Request, resp *restful.Response) {
	userId, _, err := getIdAndAuthor(req)
	if err != nil {
		fmt.Println(err)
		io.WriteString(resp.ResponseWriter, "error")
	}

	row, err := DB.Query("SELECT uid, uname, uauthority, ucreatetime, ucreator from users WHERE uid=?", userId)
	if err != nil {
		fmt.Println(err)
		io.WriteString(resp.ResponseWriter, "error")
		return
	}

	var us user
	for row.Next() {
		row.Scan(&us.uid, &us.uname, &us.uauthority, &us.ucreatetime, &us.ucreator)
	}

	m := make(map[string]interface{})

	usertomap(&us, m)
	fmt.Println(m)
	resp.WriteEntity(m)
}

//resuful api, delete user that has low authority
//success return "delete success"
//error return "error"
func deleteUser(req *restful.Request, resp *restful.Response) {
	userId, authority, err := getIdAndAuthor(req)
	if err != nil {
		fmt.Println(err)
		io.WriteString(resp.ResponseWriter, "error")
		return
	}
	fmt.Println(userId, authority)

	row, err := DB.Query("SELECT uauthority from users WHERE uid=?", userId)
	if err != nil {
		fmt.Println(err)
		io.WriteString(resp.ResponseWriter, "error")
		return
	}
	var uauthority int
	for row.Next() {
		row.Scan(&uauthority)
	}

	if authority >= uauthority {
		fmt.Println("low authority")
		io.WriteString(resp.ResponseWriter, "error")
		return
	}

	_, err = DB.Exec("DELETE from users WHERE uid=?", userId)
	if err != nil {
		fmt.Println("delete error")
		io.WriteString(resp.ResponseWriter, "error")
		return
	}

	resp.AddHeader("Content-Type", "text/plain")
	io.WriteString(resp.ResponseWriter, "delete success")
	//connect db

}

//resuful api, update user that has low authority
//success return "update user success"
//error return "error"
func putUser(req *restful.Request, resp *restful.Response) {
	userId, authority, err := getIdAndAuthor(req)
	if err != nil {
		fmt.Println(err)
		io.WriteString(resp.ResponseWriter, "error")
		return
	}

	row, err := DB.Query("SELECT uauthority from users WHERE uid=?", userId)
	if err != nil {
		fmt.Println(err)
		io.WriteString(resp.ResponseWriter, "error")
		return
	}
	var uauthority int
	for row.Next() {
		row.Scan(&uauthority)
	}

	if authority >= uauthority {
		fmt.Println("low authority")
		io.WriteString(resp.ResponseWriter, "error")
		return
	}

	doc := make(map[string]interface{})
	err = req.ReadEntity(&doc)
	if err != nil {
		fmt.Println(err)
		io.WriteString(resp.ResponseWriter, "error")
		return
	}
	fmt.Println(doc)

	var str []byte
	for k, v := range doc {
		if newv, ok := v.(string); ok {
			tmp := fmt.Sprintf("%s='%s'", k, newv)
			str = append(str, tmp...)
		} else {
			tmp := fmt.Sprintf("%s=%s", k, v)
			str = append(str, tmp...)
		}
	}

	stmt := fmt.Sprintf("update users set %s where uid = %d", str, userId)
	fmt.Println(stmt)
	_, err = DB.Exec(stmt)
	if err != nil {
		fmt.Println(err)
		io.WriteString(resp.ResponseWriter, "error")
		return
	}

	resp.AddHeader("Content-Type", "text/plain")
	io.WriteString(resp.ResponseWriter, "update user success")
}

//resuful api, add new user that has low authority
//success return "add new user success"
//error return "error"
func postUser(req *restful.Request, resp *restful.Response) {
	doc := make(map[string]interface{})
	err := req.ReadEntity(&doc)
	if err != nil {
		fmt.Println(err)
		io.WriteString(resp.ResponseWriter, "error")
		return
	}
	fmt.Println(doc)

	var col, val []byte
	for k, v := range doc {
		col = append(col, fmt.Sprintf("%s,", k)...)
		if newv, ok := v.(string); ok {
			val = append(val, fmt.Sprintf("'%s',", newv)...)
		} else {
			val = append(val, fmt.Sprintf("%s,", v)...)
		}

	}

	stmt := fmt.Sprintf("INSERT INTO users(%s) VALUES(%s)", col[:len(col)-1], val[:len(val)-1])
	fmt.Println(stmt)
	_, err = DB.Exec(stmt)
	if err != nil {
		fmt.Println(err)
		io.WriteString(resp.ResponseWriter, "error")
		return
	}

	resp.AddHeader("Content-Type", "text/plain")
	io.WriteString(resp.ResponseWriter, "add new user success")
}

//resuful api, get all users that lower or equal authority
//success return json
//error return "error"
func getUsers(req *restful.Request, resp *restful.Response) {
	authority, err := strconv.Atoi(req.QueryParameter("authority"))
	if err != nil {
		fmt.Println(err)
		io.WriteString(resp.ResponseWriter, "error")
		return
	}
	fmt.Println(authority)

	row, err := DB.Query("SELECT uid, uname, uauthority, ucreatetime, ucreator from users WHERE uauthority >= ?", authority)
	if err != nil {
		fmt.Println(err)
		io.WriteString(resp.ResponseWriter, "error")
		return
	}

	var us user
	m := make(map[string]interface{})
	var re []interface{}

	for row.Next() {
		row.Scan(&us.uid, &us.uname, &us.uauthority, &us.ucreatetime, &us.ucreator)
		usertomap(&us, m)
		re = append(re, m)
	}
	resp.WriteEntity(re)
}

func UserSrvConfig() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/user").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/{user-id}").Param(ws.QueryParameter("authority", "authority").DataType("int")).To(getUser))
	ws.Route(ws.DELETE("/{user-id}").Param(ws.QueryParameter("authority", "authority").DataType("int")).To(deleteUser))
	ws.Route(ws.PUT("/{user-id}").Param(ws.QueryParameter("authority", "authority").DataType("int")).To(putUser))
	ws.Route(ws.POST("").To(postUser))
	ws.Route(ws.GET("").Param(ws.QueryParameter("authority", "authority").DataType("int")).To(getUsers))

	return ws
}
