package report

import (
	"github.com/hongxufeng/fileLogger"
	"utils/config"
	"utils/service"
	"model"
	"datahelper/report"
	"fmt"
)

var Info *fileLogger.FileLogger
var Error *fileLogger.FileLogger
type ReportModule struct {
	level service.LEVEL
}

func (module *ReportModule) Init(conf *config.Config) error {
	module.level=service.SetEnvironment(conf.Environment)
	Info=fileLogger.NewDefaultLogger(conf.LogDir, "Report_Info.log")
	Error=fileLogger.NewDefaultLogger(conf.LogDir, "Report_Error.log")
	Info.SetPrefix("[REPORT] ")
	Error.SetPrefix("[REPORT] ")
	return nil
}

func (module *ReportModule) User_GetTable(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var settings model.Settings
	err = req.GetParams("table", &settings.TableID, "page", &settings.Page, "rows", &settings.Rows, "colpage", &settings.ColPage)
	if err != nil {
		return
	}
	err = req.ParseEncodeUrl("configFile", &settings.ConfigFile, "hasCheckbox", &settings.HasCheckbox, "style", &settings.Style, "rowList", &settings.RowList)
	if err != nil {
		return
	}
	if module.level>=service.DEV{
		fmt.Println(settings)
	}
	param, err := report.New(req.Uid, settings)
	if (err != nil) {
		return
	} else {
		result["res"], err = param.GetTable()
	}
	return
}
func (module *ReportModule) User_SearchTree(req *service.HttpRequest, result map[string]interface{}) (err error) {
	return
}

func (module *ReportModule) User_LocateNode(req *service.HttpRequest, result map[string]interface{}) (err error) {
	return
}