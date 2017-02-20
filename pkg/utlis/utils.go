package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/golang/glog"
)

func GetStringEnv(name, def string) string {
	val := os.Getenv(name)
	if val == "" {
		fmt.Printf("%s use default value %s\n", name, def)
		return def
	}
	fmt.Printf("%s use value %s\n", name, val)
	return val
}

func GetNowTime() time.Time {
	return time.Now()
}

func GetPosition() (string, string, int) {
	funcName, file, line, ok := runtime.Caller(1)
	if ok {
		return file, runtime.FuncForPC(funcName).Name(), line
	}
	return "", "", 0
}

func Readresp(resp *http.Response) []byte {
	bodystr := make([]byte, 0)
	buf := make([]byte, 512)
	nn, err := resp.Body.Read(buf)
	for err == nil {
		bodystr = append(bodystr, buf[:nn]...)
		nn, err = resp.Body.Read(buf)
	}
	if err == io.EOF {
		bodystr = append(bodystr, buf[:nn]...)
	} else {
		file, funcn, line := GetPosition()
		glog.Errorf("exec readresp with error %v, at %s %s %d\n", err, file, funcn, line)
	}

	return bodystr
}
