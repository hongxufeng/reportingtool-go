package user

import (
	"github.com/aiwuTech/fileLogger"
	"utils/config"
)

type UserModule struct {
	log *fileLogger.FileLogger
}

func (module *UserModule) Init(conf *config.Config) error {
	module.log=fileLogger.NewDefaultLogger(conf.LogDir, "User.log")
	return nil
}

