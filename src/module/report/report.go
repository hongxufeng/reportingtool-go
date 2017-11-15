package report

import (
	"github.com/hongxufeng/fileLogger"
	"utils/config"
	"utils/service"
	"model"
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

func (module *ReportModule) User_Reportingtool (req *service.HttpRequest, result map[string]interface{}) (e error) {
	var param  model.Param
	e=req.ParseEncodeUrl("cmd",&param.Settings.Cmd)
	if e!=nil {
		return
	}
	if param.Settings.Cmd==model.CMD_GetTable{
		e=req.GetParams("table",&param.Settings.TableID,"page",&param.Settings.Page,"rows",&param.Settings.Rows,"colpage",&param.Settings.ColPage)
		if e!=nil {
			return
		}
		e=req.ParseEncodeUrl("configFile",&param.Settings.ConfigFile,"hasCheckbox",&param.Settings.HasCheckbox,"style",&param.Settings.Style,"rowList",&param.Settings.RowList)
		if e!=nil{
			return
		}
		result["res"]=param.Settings.RowList
	}

	return
}

