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
)

type Server struct {
	Info *fileLogger.FileLogger
	Error *fileLogger.FileLogger
	mvaliduser   func(r *http.Request, bodybytes []byte) (uid uint32, md5ok bool) //加密方式    如果不是合法用户，需要返回0
}

func New() (server Server, err error) {
	server.Info=fileLogger.NewDefaultLogger("/log", "info.log")
	server.Info.SetPrefix("[INFO] ")
	server.Error=fileLogger.NewDefaultLogger("/log", "error.log")
	server.Error.SetPrefix("[ERROR] ")
	server.mvaliduser=mValidUser
	return
}

func mValidUser(r *http.Request, bodybytes []byte) (uid uint32, md5ok bool) {
	md5ok = true
	c, e := r.Cookie("auth")
	if e != nil {
		return 0, md5ok
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
		return 0, md5ok
	}
	return uid, md5ok
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
		server.Error.Println("服务启动错误：%s",err)
	}else {
		server.Error.Println("http服务启动！")
	}
	return err
}

func (server *Server) BaseHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now().UnixNano()
	var result map[string]interface{} = make(map[string]interface{})
	var err error
	//bodyBytes, e := ioutil.ReadAll(r.Body)
	//if e != nil {
	//	err = NewError(ERR_INTERNAL, "read http data error : "+e.Error())
	//}
	//var body []byte
	//
	//fields := strings.Split(r.URL.Path[1:], "/")
	//if len(fields) >= 3 {
	//	body, err = handleRequest(fields[1], "X_"+fields[2], uid, r, result, bts)
	//} else {
	//	err = NewError(ERR_INVALID_PARAM, "invalid url format : "+r.URL.Path)
	//}
	//end := time.Now().UnixNano()
	//server.processErrorX(w, r, err, body, result, end-start)
}


func (server *Server) UserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "您要创建 %s!\n")
}
