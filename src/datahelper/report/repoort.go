package report

import (
	"model"
)

type Param struct {
	XmlTable interface{}//XML获得的数据结构
	Settings model.Settings
	Uid uint32
	IsAdmin bool  //用户判断得到的权限
	ColConfigDict []model.ColumnConfig
}

func New(uid uint32,settings model.Settings) (param *Param,err error){
	var XmlTable interface{}
	var ColConfigDict []model.ColumnConfig
	param=&Param{XmlTable,settings,uid,true,ColConfigDict}
	return
}
func (param *Param) GetTable() (res map[string]interface{},err error){
	res=make(map[string]interface{}, 0)
	res["search"]="search"
	res["body"]="body"
	res["selector"]="selector"
	res["condition"]="condition"
	res["row"]="row"
	return
}
func (param *Param) SearchTree() (res map[string]interface{},err error){
	return
}
func (param *Param) LocateNode() (res map[string]interface{},err error){
	return
}