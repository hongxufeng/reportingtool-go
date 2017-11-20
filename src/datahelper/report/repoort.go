package report

import (
	"model"
	"github.com/beevik/etree"
	"fmt"
	"datahelper/db"
)

type Param struct {
	Settings model.Settings
	Uid uint32
	Power int8  //用户判断得到的权限
	ColConfigDict []model.ColumnConfig
}

func New(uid uint32,settings model.Settings) (param *Param,err error){
	var ColConfigDict []model.ColumnConfig
	doc := etree.NewDocument()
	filename:="xml/"+settings.ConfigFile+".xml"
	//fmt.Println(filename)
	err = doc.ReadFromFile(filename)
	if err != nil {
		return
	}
	path:="./tables/table[@id='"+settings.TableID+"']/*";
	//fmt.Println(path)
	for _, e := range doc.FindElements(path) {
		fmt.Printf("%s: %s\n", e.Tag, e.Text())
		//赋值ColConfigDict
	}
	//根据uid判断权限
	ud,err:=db.GetUserInfo(uid)
	param=&Param{settings,uid,ud.Power,ColConfigDict}
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