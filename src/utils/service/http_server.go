package service

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"io/ioutil"
	"strings"
	"fmt"
	"github.com/aiwuTech/fileLogger"
)

type Server struct {
	log *fileLogger.FileLogger
	mvaliduser   func(r *http.Request, bodybytes []byte) (uid uint32, md5ok bool) //加密方式    如果不是合法用户，需要返回0
}

func New() (server *Server, err error) {
	server.log=fileLogger.NewDefaultLogger("/log", "log.log")
	server.log.SetLogLevel(fileLogger.INFO)
	return
}
func (server Server) StartService() error {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/ReportingTool", FormBuild)
	r.HandleFunc("/UserLogin", server.LoginHandler)
	r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("web/"))))

	// Bind to a port and pass our router in
	err :=http.ListenAndServe(":8080", r)
	if err!=nil {
		server.log.Error("服务启动错误：%s",err)
	}else {
		server.log.Println("http服务启动！")
	}
	return err
}

func (server *Server) SetmValid(f func(r *http.Request, bodybytes []byte) (uid uint32, md5ok bool)) {
	server.mvaliduser = f
}

func (server Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	//start := time.Now().UnixNano()
	//var result map[string]interface{} = make(map[string]interface{})
	//var err error
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


func FormBuild(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "您要创建 %s!\n")
}
