package service

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"github.com/aiwuTech/fileLogger"
	"strings"
	"utils"
	"datahelper"
	"time"
	"encoding/json"
)

type Server struct {
	Info *fileLogger.FileLogger
	Error *fileLogger.FileLogger
	mvaliduser   func(r *http.Request) (uid uint32) //加密方式    如果不是合法用户，需要返回0
}

func New() (server Server, err error) {
	server.Info=fileLogger.NewDefaultLogger("/log", "info.log")
	server.Info.SetPrefix("[INFO] ")
	server.Error=fileLogger.NewDefaultLogger("/log", "error.log")
	server.Error.SetPrefix("[ERROR] ")
	server.mvaliduser=mValidUser
	return
}

func mValidUser(r *http.Request) (uid uint32) {
	c, e := r.Cookie("auth")
	if e != nil {
		return 0
	}
	auth := c.Value
	var hashcode string
	var ks []string
	if strings.Contains(auth, "%09") {
		ks = strings.Split(auth, "%09")
	} else {
		ks = strings.Split(auth, "_")
	}

	if len(ks) == 2 {
		uid, e = utils.ToUint32(ks[0])
		if e != nil {
			fmt.Println(e.Error())
		}
		hashcode = ks[1]
	}
	valid, e := datahelper.UserValid(uid, hashcode)
	if e != nil || !valid {
		return 0
	}
	return uid
}

func (server Server) StartService() error {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/user", server.UserHandler)
	r.HandleFunc("/base", server.BaseHandler)
	r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("web/"))))

	// Bind to a port and pass our router in
	err :=http.ListenAndServe(":8080", r)
	if err!=nil {
		server.Error.Print("服务启动错误：%s",err)
	}else {
		server.Info.Print("http服务启动！")
	}
	return err
}

func (server *Server) UserHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now().UnixNano()
	var result map[string]interface{} = make(map[string]interface{})
	var err error
	var body []byte
	var uid uint32
	uid= server.mvaliduser(r)
	if uid > 0 {
		fields := strings.Split(r.URL.Path[1:], "/")
		if len(fields) >= 3 {
			body, err = server.RequestHandler(fields[1], "User_"+fields[2], uid, r, result)
		} else {
			err = NewError(ERR_INVALID_PARAM, "invalid url format : "+r.URL.Path)
		}
	} else {
		err = NewError(ERR_INVALID_USER, "invalid user")
	}
	end := time.Now().UnixNano()
	server.Respose(w, r, err, body, result, end-start)
}

func (server *Server) BaseHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now().UnixNano()
	var result map[string]interface{} = make(map[string]interface{})
	var err error
	var body []byte
	var uid uint32
	uid = server.mvaliduser(r)
	fields := strings.Split(r.URL.Path[1:], "/")
	if len(fields) >= 3 {
		body, err = server.RequestHandler(fields[1], "Base_"+fields[2], uid, r, result)
	} else {
		err = NewError(ERR_INVALID_PARAM, "invalid url format : "+r.URL.Path)
	}
	end := time.Now().UnixNano()
	server.Respose(w, r, err, body, result, end-start)
}

func (server *Server) RequestHandler(moduleName string, methodName string, uid uint32, r *http.Request, result map[string]interface{}) (byte []byte,e error) {
	return
}

func (server *Server) Respose(w http.ResponseWriter, r *http.Request, err error, reqBody []byte, result map[string]interface{}, duration int64) {
	var re Error
	switch e := err.(type) {
	case nil:
	case Error:
		re = e
	default:
		re = NewError(ERR_INTERNAL, e.Error(), "未知错误")
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	if re.Code == ERR_NOERR {
		result["status"] = "ok"
		result["tm"] = time.Now().UnixNano()
		res, e := json.Marshal(result)
		if e == nil {
			server.Log(r, w, reqBody, []byte(res), true, duration)
		} else {
			server.ResposeErr(r, w, reqBody, NewError(ERR_INTERNAL, fmt.Sprintf("Marshal result error : %v", e.Error())), duration)
		}
	} else {
		server.ResposeErr(r, w, reqBody, re, duration)
	}
}
func (server *Server) ResposeErr(r *http.Request, w http.ResponseWriter, reqBody []byte, err Error, duration int64) {
	var result map[string]interface{} = make(map[string]interface{})
	result["status"] = "fail"
	result["code"] = err.Code
	result["msg"] = err.Show
	result["detail"] = err.Desc
	res, _ := json.Marshal(result)
	server.Log(r, w, reqBody, res, false, duration)
}
func (server *Server) Log(r *http.Request, w http.ResponseWriter, reqBody []byte, result []byte, success bool, duration int64) {
	w.Write(result)
	var l string
	var uid, response string
	uidCookie, e := r.Cookie("auth")
	if e != nil {
		uid = "无"
	} else {
		uid = uidCookie.Value
	}
	if reqBody != nil {
		response = string(reqBody)
	}
	format := " duration: %.2fms |"
	format += "uid: %s  %s|"
	format += "uri: %s |"
	format += "param: %s |"
	format += "response: %s |"
	addr := strings.Split(r.RemoteAddr, ":")
	l = fmt.Sprintf(format, float64(duration)/1000000, uid, addr[0], r.URL.String(), response, string(result))
	if success {
		server.Info.Print(l)
	}else {
		server.Error.Print(l)
	}
}