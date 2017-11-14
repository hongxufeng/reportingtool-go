package report

import (
	"github.com/hongxufeng/fileLogger"
	"utils/config"
)

type ReportModule struct {
	log *fileLogger.FileLogger
}

func (module *ReportModule) Init(conf *config.Config) error {
	module.log=fileLogger.NewDefaultLogger(conf.LogDir, "Report.log")
	return nil
}

