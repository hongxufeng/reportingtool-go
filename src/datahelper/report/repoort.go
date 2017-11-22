package report

import (
	"model"
	"github.com/beevik/etree"
	"fmt"
	"datahelper/db"
	"strings"
	"utils/service"
	"bytes"
	"utils/function"
)

type Param struct {
	TableConfig model.TableConfig
	Settings model.Settings
	Uid uint32
	Power uint8  //用户判断得到的权限 暂时用0代表最高权限
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
	tableconfig:=model.TableConfig{}
	if tableelement:=doc.FindElements("./tables/table[@id='"+settings.TableID+"']");len(tableelement)>0{
		fmt.Println(len(tableelement))
		table:=tableelement[0]
		defaultorder := table.SelectAttr("defaultorder")
		if defaultorder!=nil{
			tableconfig.HasDefaultOrder=true
			tableconfig.DefaultOrder=defaultorder.Value
		}
		name := table.SelectAttr("name")
		if name==nil{
			err=service.NewError(service.ERR_XML_ATTRIBUTE_LACK,"XML配置属性name缺失哦！")
		}else{
			tableconfig.Name=name.Value
		}
		adminname := table.SelectAttr("adminname")
		if adminname!=nil{
			tableconfig.HasAdminName=true
			tableconfig.AdminName=adminname.Value
		}
		power := table.SelectAttr("power")
		if power!=nil{
			tableconfig.HasPower=true
			tableconfig.Power,_=function.StringToUint8(power.Value)
		}
		excel := table.SelectAttr("excel")
		if excel!=nil{
			tableconfig.HasExcel=true
			tableconfig.Excel=excel.Value
		}
	}else {
		err=service.NewError(service.ERR_XML_TABLE_LACK,"您配置的XML表格是不在的啊！")
	}
	path:="./tables/table[@id='"+settings.TableID+"']/*";
	//fmt.Println(path)
	for _, elemnt := range doc.FindElements(path) {
		fmt.Printf("%s: %s\n", elemnt.Tag, elemnt.Text())
		//赋值ColConfigDict
		cc:=model.ColumnConfig{}
		cc.Tag=elemnt.Tag
		cc.Text=elemnt.Text()
		btnicon := elemnt.SelectAttr("btn-icon")
		btnfunc := elemnt.SelectAttr("btn-func")
		if btnicon!=nil&&btnfunc!=nil {
			cc.HasBtn=true
			cc.BtnFunc=btnfunc.Value
			cc.BtnIcon=btnicon.Value
		}
		dateformat := elemnt.SelectAttr("dateformat")
		if dateformat!=nil{
			cc.HasDateformat=true
			cc.Dateformat=dateformat.Value
		}
		checkbox := elemnt.SelectAttr("checkbox")
		if checkbox!=nil{
			cc.HasCheckBox=true
			cc.CheckBox=checkbox.Value
		}
		defaultvalue := elemnt.SelectAttr("defaultvalue")
		if defaultvalue!=nil{
			cc.HasDefaultValue=true
			cc.DefaultValue=defaultvalue.Value
		}
		formatter := elemnt.SelectAttr("formatter")
		if formatter!=nil{
			cc.HasFormatter=true
			cc.Formatter=formatter.Value
		}
		selector := elemnt.SelectAttr("selector")
		if selector!=nil{
			cc.IsInselector=true
			cc.Selector=selector.Value
		}
		formatterr := elemnt.SelectAttr("formatter-r")
		if formatterr!=nil{
			cc.HasFormatterR=true
			cc.FormatterR=formatterr.Value
		}
		searchtype := elemnt.SelectAttr("search-type")
		if searchtype!=nil{
			cc.HasSearchType=true
			cc.SearchType=searchtype.Value
		}
		selectorfunc := elemnt.SelectAttr("selector-func")
		if selectorfunc!=nil{
			cc.IsInselector=true
			cc.SelectorFunc=selectorfunc.Value
		}
		selectortext := elemnt.SelectAttr("selector-text")
		if selectortext!=nil{
			cc.SelectorText=selectortext.Value
		}
		linkto := elemnt.SelectAttr("linkto")
		passedcol := elemnt.SelectAttr("passedcol")
		if linkto!=nil&&passedcol!=nil{
			cc.HasLinkTo=true
			cc.LinkTo=linkto.Value
			cc.Passedcol =strings.Split(passedcol.Value,",")
			ignoredpassedcol := elemnt.SelectAttr("ignoredpassedcol")
			if param.Power==0&&ignoredpassedcol!=nil{
				ipd:=strings.Split(ignoredpassedcol.Value,",")
				for  no,_:=range ipd{
					cc.Passedcol=append(cc.Passedcol[:no],cc.Passedcol[no+1:]...)
				}
			}
		}
		selectormulti := elemnt.SelectAttr("selector-multi")
		if selectormulti!=nil&&selectormulti.Value=="true"{
			cc.HasSelectorMulti=true
		}
		navname := elemnt.SelectAttr("navname")
		if navname!=nil{
			cc.HasNavname=true
			cc.Navname=navname.Value
		}
		search4admin := elemnt.SelectAttr("search4admin")
		if search4admin!=nil&&search4admin.Value=="true"{
			cc.Search4Admin=true
		}
		timetransfer := elemnt.SelectAttr("timetransfer")
		if timetransfer!=nil{
			cc.HasTimetransfer=true
			cc.Timetransfer=timetransfer.Value
		}
		precision := elemnt.SelectAttr("precision")
		if precision!=nil{
			cc.HasPrecision=true
			cc.Precision=precision.Value
		}
		visibility := elemnt.SelectAttr("visibility")
		if visibility!=nil{
			cc.Visibility=visibility.Value
		}
		percentageform := elemnt.SelectAttr("percentageform")
		if percentageform!=nil&&percentageform.Value=="true"{
			cc.IsInPercentageform=true
		}
		ColConfigDict=append(ColConfigDict, cc)
	}
	fmt.Println(ColConfigDict)
	fmt.Println(tableconfig)
	//根据uid判断权限
	ud,err:=db.GetUserInfo(uid)
	param=&Param{tableconfig,settings,uid,ud.Power,ColConfigDict}
	return
}
func BuildSQL(param *Param) (string,error){
	var buf bytes.Buffer
	if param.TableConfig.HasAdminName&&param.Power==0 {
		buf.WriteString(param.TableConfig.Name)
	}else {
		buf.WriteString(param.TableConfig.Name)
	}
	if param.TableConfig.HasDefaultOrder{
		buf.WriteString("order by ")
		buf.WriteString(param.TableConfig.DefaultOrder)
	}
	if param.TableConfig.HasPower&&param.Power>=param.TableConfig.Power {
		return buf.String(),service.NewError(service.ERR_POWER_DENIED,"您的用户权限不足啊！")
	}
	return buf.String(),nil
}
func (param *Param) GetTable() (res map[string]interface{},err error){
	res=make(map[string]interface{}, 0)
	query,err:=BuildSQL(param)
	if err!=nil{
		return
	}
	fmt.Println(query)
	res["search"]=GetTableSearch(param)
	res["body"]=GetTableBody(param)
	res["selector"]=GetTableSelector(param)
	res["condition"]=GetTableCondition(param)
	res["row"]=GetTableRow(param)
	return
}
func (param *Param) SearchTree() (res map[string]interface{},err error){
	return
}
func (param *Param) LocateNode() (res map[string]interface{},err error){
	return
}