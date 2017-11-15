package user


import (
	"utils/config"
	"utils/service"
	"datahelper/user"
	"github.com/hongxufeng/fileLogger"
	"model"
)

type UserModule struct {
	info *fileLogger.FileLogger
	error *fileLogger.FileLogger
}

func (module *UserModule) Init(conf *config.Config) error {
	module=&UserModule{fileLogger.NewDefaultLogger(conf.LogDir, "User_Info.log"),fileLogger.NewDefaultLogger(conf.LogDir, "User_Error.log")}
	module.info.SetPrefix("[SERVICE] ")
	module.error.SetPrefix("[SERVICE] ")
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
		module.error.Error("%s  auth failed !",loginData.Username)
		//这里增加判断登录频繁次数
		result["res"] = user.CreateFailResp( service.ERR_INVALID_USER, "少侠，您输入的密码有误啊！")
		return
	}
	result["res"]= user.CreateSuccessResp(ud)
	return
}