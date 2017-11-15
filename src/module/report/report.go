package report

import (
	"github.com/hongxufeng/fileLogger"
	"utils/config"
	"utils/service"
	"model"
	"datahelper/report"
)

type ReportModule struct {
	info  *fileLogger.FileLogger
	error *fileLogger.FileLogger
}

func (module *ReportModule) Init(conf *config.Config) error {
	module = &ReportModule{fileLogger.NewDefaultLogger(conf.LogDir, "Report_Info.log"), fileLogger.NewDefaultLogger(conf.LogDir, "Report_Error.log")}
	module.info.SetPrefix("[SERVICE] ")
	module.error.SetPrefix("[SERVICE] ")
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