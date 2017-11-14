package report

import (
	"github.com/hongxufeng/fileLogger"
	"utils/config"
)

type ReportModule struct {
	info *fileLogger.FileLogger
	error *fileLogger.FileLogger
}

func (module *ReportModule) Init(conf *config.Config) error {
	module=&ReportModule{fileLogger.NewDefaultLogger(conf.LogDir, "Report_Info.log"),fileLogger.NewDefaultLogger(conf.LogDir, "Report_Error.log")}
	module.info.SetPrefix("[SERVICE] ")
	module.error.SetPrefix("[SERVICE] ")
	return nil
}

