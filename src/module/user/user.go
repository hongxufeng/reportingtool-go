package user

import (
	"github.com/aiwuTech/fileLogger"
	"utils/config"
	"utils/service"
)

type UserModule struct {
	log *fileLogger.FileLogger
}

func (module *UserModule) Init(conf *config.Config) error {
	module.log=fileLogger.NewDefaultLogger(conf.LogDir, "User.log")
	return nil
}

func (module *UserModule) Base_UserLogin(req *service.HttpRequest, result map[string]interface{}) (e error) {

	return nil
}