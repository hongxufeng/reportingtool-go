package service

type DefaultModule struct {
}

func (module *DefaultModule) Init() error {
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