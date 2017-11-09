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
	"reflect"
	"io/ioutil"
	"utils/config"
)

type Server struct {
	modules      map[string]Module
	service *fileLogger.FileLogger
	conf         *config.Config
	mvaliduser   func(r *http.Request) (uid uint32,err error) //加密方式    如果不是合法用户，需要返回0
	parseBody    bool                               //是否把POST的内容解析为json对象
	customResult bool                               //返回结果中是否包含result和tm项
	multipart    bool                               //是否multipart post 上传文件
}

func New(conf *config.Config,args ...bool) (server *Server, err error) {
	server = &Server{make(map[string]Module), fileLogger.NewDefaultLogger(conf.LogDir, "Service.log"), conf, mValidUser, true, false, false}
	server.service.SetPrefix("[SERVICE] ")
	if len(args) >= 1 {
		server.parseBody = args[0]
	}
	if len(args) >= 2 {
		server.customResult = args[1]
	}
	if len(args) >= 3 {
		server.multipart = args[2]
	}
	err=server.AddModule("default", &DefaultModule{})
	return
}

func mValidUser(r *http.Request) (uid uint32,err error) {
	c, e := r.Cookie("auth")
	if e != nil {
		return 0,nil
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
		return 0,e
	}
	return uid,nil
}

func (server *Server) AddModule(name string, module Module) (err error) {
	fmt.Printf("add module %s... ", name)
	err = module.Init(server.conf)
	if err != nil {
		return
	}
	fmt.Println("ok")
	server.service.Printf("add module %s success",name)
	server.modules[name] = module
	return
}

func (server *Server) StartService() error {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/user/{module}/{method}", server.UserHandler)
	r.HandleFunc("/base/{module}/{method}", server.BaseHandler)
	r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("web/"))))
	fmt.Print("服务已经启动！")
	server.service.Print("服务已经启动！")
	// Bind to a port and pass our router in
	return http.ListenAndServe(server.conf.Address.Port, r)
}

func (server *Server) UserHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now().UnixNano()
	var result map[string]interface{} = make(map[string]interface{})
	var err error
	var body []byte
	var uid uint32
	vars := mux.Vars(r)
	uid,err= server.mvaliduser(r)
	if err==nil {
		if uid > 0 {
			body, err = server.RequestHandler(vars["module"], "User_"+vars["method"], uid, r, result, nil)
		} else {
			err = NewError(ERR_INVALID_USER, "invalid user")
		}
	}
	end := time.Now().UnixNano()
	server.Respose(w, r, err, body, result, end-start)
}

func (server *Server) BaseHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now().UnixNano()
	vars := mux.Vars(r)
	var result map[string]interface{} = make(map[string]interface{})
	var err error
	var body []byte
	var uid uint32
	uid,err = server.mvaliduser(r)
	if err==nil {
		body, err = server.RequestHandler(vars["module"], "Base_"+vars["method"], uid, r, result, nil)
	}
	end := time.Now().UnixNano()
	server.Respose(w, r, err, body, result, end-start)
}

//如果参数bodyBytes不是空，则可代表解密后的Request.body的内容
func (server *Server) RequestHandler(moduleName string, methodName string, uid uint32, r *http.Request, result map[string]interface{},bodyBytes []byte) (byte []byte,e error) {
	if server.multipart == false {
		if bodyBytes == nil {
			var e error
			bodyBytes, e = ioutil.ReadAll(r.Body)
			if e != nil {
				return nil, NewError(ERR_INTERNAL, "read http data error : "+e.Error())
			}
		}
	} else {
		bodyBytes = nil
	}
	var body map[string]interface{}
	 //rr, e := r.MultipartReader()
	 //fmt.Println(fmt.Sprintf("r.MultipartReader rr %v| e %v|bodyBytes %v|r.MultipartForm %v  ", rr, e, bodyBytes, r.MultipartForm))
	if len(bodyBytes) == 0 {
		//get请求
		body = make(map[string]interface{})
	} else if server.parseBody {
		e := json.Unmarshal(bodyBytes, &body)
		if e != nil {
			return bodyBytes, NewError(ERR_INVALID_PARAM, "read body error : "+e.Error())
		}
	}
	var values []reflect.Value
	module, ok := server.modules[moduleName]
	if ok {
		method := reflect.ValueOf(module).MethodByName(methodName)
		if method.IsValid() {
			values = method.Call([]reflect.Value{reflect.ValueOf(&HttpRequest{body, bodyBytes, r, uid}), reflect.ValueOf(result)})
		} else {
			method = reflect.ValueOf(server.modules["default"]).MethodByName("ErrorMethod")
			values = method.Call([]reflect.Value{reflect.ValueOf(&HttpRequest{body, bodyBytes, r, uid}), reflect.ValueOf(result)})
		}
	} else {
		method := reflect.ValueOf(server.modules["default"]).MethodByName("ErrorModule")
		values = method.Call([]reflect.Value{reflect.ValueOf(&HttpRequest{body, bodyBytes, r, uid}), reflect.ValueOf(result)})
	}
	if len(values) != 1 {
		return bodyBytes, NewError(ERR_INTERNAL, fmt.Sprintf("method %s.%s return value is not right.", moduleName, methodName))
	}
	switch x := values[0].Interface().(type) {
	case nil:
		return bodyBytes, nil
	default:
		return bodyBytes, x.(error)
	}
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
		if !server.customResult {
			result["status"] = "ok"
			result["tm"] = time.Now().UnixNano()
		}
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
	if success {
		server.service.Info(format, float64(duration)/1000000, uid, addr[0], r.URL.String(), response, string(result))
	}else {
		server.service.Error(format, float64(duration)/1000000, uid, addr[0], r.URL.String(), response, string(result))
	}
}