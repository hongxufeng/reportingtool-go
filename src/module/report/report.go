package report

import (
	"github.com/hongxufeng/fileLogger"
	"utils/config"
	"utils/service"
	"model"
	"datahelper/report"
	"errors"
	"fmt"
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

func (module *ReportModule) User_Reportingtool (req *service.HttpRequest, result map[string]interface{}) (err error) {
	var settings  model.Settings
	err=req.ParseEncodeUrl("cmd",&settings.Cmd)
	if err!=nil {
		return
	}
	if settings.Cmd==model.CMD_GetTable{
		err=req.GetParams("table",&settings.TableID,"page",&settings.Page,"rows",&settings.Rows,"colpage",&settings.ColPage)
		if err!=nil {
			return
		}
		err=req.ParseEncodeUrl("configFile",&settings.ConfigFile,"hasCheckbox",&settings.HasCheckbox,"style",&settings.Style,"rowList",&settings.RowList)
		if err!=nil{
			return
		}
		param,e:=report.New(req.Uid,settings)
		if(e!=nil){
			return e
		}else {
			result["res"],err=param.GetTable()
		}
	}else if settings.Cmd==model.CMD_SearchTree {
		
	}else if settings.Cmd==model.CMD_LocateNode {

	}else {
		return errors.New(fmt.Sprintf("cmd [%v] is not support",settings.Cmd))
	}
	return
}

