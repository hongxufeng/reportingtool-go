package user


import (
	"github.com/aiwuTech/fileLogger"
	"utils/config"
	"utils/service"
	"datahelper/user"
	"model"
)

type UserModule struct {
	log *fileLogger.FileLogger
}

func (module *UserModule) Init(conf *config.Config) error {
	module.log=fileLogger.NewDefaultLogger(conf.LogDir, "User.log")
	return nil
}

func (module *UserModule) Base_UserLogin(req *service.HttpRequest, result map[string]interface{}) (e error) {
	param :=model.LoginData{}
	e = req.ParseAjax(&param)
	if e != nil {
		return
	}
	//判断登录频繁，防止暴力破解
	if forbid, _ := user.CheckUserForbid(param.Username);forbid {
		result["res"] = user.CreateFailResp(service.ERR_LOGIN_FREQUENT, "登陆过于频繁,请稍候再试")
		return
	}
	if state, _ := user.CheckUserState(param.Username);state {
		result["res"] = user.CreateFailResp(service.RR_STATUS_DENIED, "用户登录状态关闭")
		return
	}

	ud, e := user.CheckAuth(param.Username, param.Password)
	if e != nil {
		module.log.Error("%s  auth failed !",param.Username)
		result["res"] = user.CreateFailResp( service.ERR_INVALID_USER, "对不起，用户名或密码错误！")
		return nil
	}
	result["res"], e = user.CreateSuccessResp(ud)
	return
}