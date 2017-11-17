package user


import (
	"utils/config"
	"utils/service"
	"datahelper/user"
	"github.com/hongxufeng/fileLogger"
	"model"
	"datahelper/db"
	"fmt"
)
var Info *fileLogger.FileLogger
var Error *fileLogger.FileLogger
type UserModule struct {
	level service.LEVEL
}

func (module *UserModule) Init(conf *config.Config) error {
	module.level=service.SetEnvironment(conf.Environment)
	Info=fileLogger.NewDefaultLogger(conf.LogDir, "User_Info.log")
	Error=fileLogger.NewDefaultLogger(conf.LogDir, "User_Error.log")
	Info.SetPrefix("[USER] ")
	Error.SetPrefix("[USER] ")
	return nil
}

func (module *UserModule) Base_UserLogin(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var loginData model.LoginData
	err=req.ParseEncodeUrl("username",&loginData.Username,"password",&loginData.Password)
	if err != nil {
		return
	}
	uid,e := user.GetUidbyName(loginData.Username)
	if uid==0||e!=nil{
		result["res"] = user.CreateFailResp(service.ERR_USER_NOT_FOUND, "貌似您输入的帐号不存在！")
		return
	}
	if state, e := user.CheckUserState(uid);!state||e!=nil {
		result["res"] = user.CreateFailResp(service.RR_STATUS_DENIED, "现如今用户登录状态关闭呢！")
		return
	}
	//判断登录频繁，防止暴力破解
	if forbid, e := user.CheckUserForbid(uid);forbid ||e!=nil{
		result["res"] = user.CreateFailResp(service.ERR_LOGIN_FREQUENT, "您登录有点频繁，请稍事休息！")
		return
	}
	ud, e := user.CheckAuth(uid, loginData.Password)
	if e != nil {
		fmt.Print(uid)
		fmt.Print(fmt.Sprintf("%d  auth failed  |",uid))
		Info.Info("%d  auth failed",uid)
		if cnt, _ := db.UserLoginErrCnt(uid); cnt >= 5 {
			db.SetUserForbid(uid)
			Info.Info("forbid user :%d",uid )
		}
		//这里增加判断登录频繁次数
		result["res"] = user.CreateFailResp( service.ERR_INVALID_USER, "少侠，您输入的密码有误啊！")
		return
	}
	result["res"]= user.CreateSuccessResp(ud)
	return
}