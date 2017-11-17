package service

import (
	"utils/config"
)

type DefaultModule struct {
	level LEVEL
}

func (module *DefaultModule) Init(conf *config.Config) error {
	module.level=SetEnvironment(conf.Environment)
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