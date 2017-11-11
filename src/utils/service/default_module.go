package service

import (
	"utils/config"
	"github.com/hongxufeng/fileLogger"
)

type DefaultModule struct {
	log *fileLogger.FileLogger
}

func (module *DefaultModule) Init(conf *config.Config) error {
	module.log=fileLogger.NewDefaultLogger(conf.LogDir, "Default.log")
	return nil
}
func (module *DefaultModule) ErrorModule(req *HttpRequest, res map[string]interface{}) (e Error) {
	e.Desc = "Invalid Module Name"
	e.Code = ERR_INVALID_PARAM
	return
}
func (module *DefaultModule) ErrorMethod(req *HttpRequest, res map[string]interface{}) (e Error) {
	e.Desc = "Invalid Method Name"
	e.Code = ERR_INVALID_PARAM
	return
}