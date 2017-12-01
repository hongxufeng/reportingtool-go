package report

import (
	"bytes"
	"model"
	"utils/function"
	"strings"
	"utils/service"
	"fmt"
	"datahelper/db"
)

func AppendWhere(req *service.HttpRequest,param *Param,buf *bytes.Buffer)  {
	hasWhere := false
	for _,colconfig:=range param.ColConfigDict{
		var value string
		_=req.GetParams(colconfig.Tag,&value)
		if len(value) == 0 {
			_ = req.GetParams(colconfig.Tag+"~~", &value)
			if len(value)==0{
				continue
			}
		}
		if hasWhere==false{
			hasWhere=true
			buf.WriteString(" where ")
		}
		buf.WriteString(colconfig.Tag)
		buf.WriteString("=")
		buf.WriteString("(")
		buf.WriteString(strings.Join(strings.Split(value,",")," or "))
		buf.WriteString(")")
		buf.WriteString(" and ")
	}
	if hasWhere{
		buf.Truncate(buf.Len()-5)
	}
}

func BuildQuerySQL(req *service.HttpRequest,param *Param) (string,error){
	var buf bytes.Buffer
	buf.WriteString("select ")
	var size = len(param.ColConfigDict)
	if size==0{
		return buf.String(),service.NewError(service.ERR_XML_ATTRIBUTE_LACK,"您至少需要配置一项XML中的列属性啊！")
	}
	for i:=0; i<size;i++ {
		if param.ColConfigDict[i].Tag=="buttons"||param.ColConfigDict[i].Tag=="pagerbuttons"{
			continue
		}
		if i!=0{
			buf.WriteString(",")
		}
		buf.WriteString(param.ColConfigDict[i].Tag)
	}
	buf.WriteString(" from ")
	buf.WriteString(param.TableConfig.Name)

	AppendWhere(req,param,&buf)
	if param.Settings.Order!=""{
		buf.WriteString(" order by ")
		buf.WriteString(param.Settings.Order)
	}else if param.TableConfig.HasDefaultOrder{
		buf.WriteString(" order by ")
		buf.WriteString(param.TableConfig.DefaultOrder)
	}
	buf.WriteString(" limit ")
	buf.WriteString(function.IntToString(param.Settings.Rows*(param.Settings.Page-1)))
	buf.WriteString(",")
	buf.WriteString(function.IntToString(param.Settings.Rows*param.Settings.Page-1))
	if param.TableConfig.HasPower&&param.Power>=param.TableConfig.Power {
		return buf.String(),service.NewError(service.ERR_POWER_DENIED,"您的用户权限不足啊！")
	}
	return buf.String(),nil
}

func GetSelectQuery(param *Param,fields string) (query string,err error){
	var buf bytes.Buffer
	buf.WriteString("select ")
	buf.WriteString(fields)
	buf.WriteString(" from ")
	buf.WriteString(param.TableConfig.Name)
	query=buf.String()
	fmt.Println(query)
	return
}

func GetTableCount(param *Param,fields string) (count int,err error){
	query,_:=GetSelectQuery(param,"count("+fields+")")
	result,err:=db.Query(query)
	if err!=nil{
		return
	}
	if result.Next(){
		err=result.Scan(&count)
	}else {
		err=service.NewError(service.ERR_MYSQL)
	}
	return
}


func BuildTablePager(param *Param,bodybuf *bytes.Buffer,count int,style string) (err error) {
	var x int
	if count % param.Settings.Rows == 0 {
		x=0
	}else {
		x=1
	}
	totalpages := count / param.Settings.Rows + x
	rowlist:=strings.Split(param.Settings.RowList,",")
	start:=(param.Settings.Page - 1) * param.Settings.Rows + 1
	var end int
	if (param.Settings.Page * param.Settings.Rows) <= count {
		end=param.Settings.Page*param.Settings.Rows
	}else {
		end=count
	}
	bodybuf.WriteString("<div class=\"rt-pager-container\">")
	bodybuf.WriteString("<div class=\"rt-pager-buttons\">")
	if param.TableConfig.HasExcel&&param.TableConfig.Excel=="true"{
		bodybuf.WriteString("<span class=\" rt-pager-export rt-pager-btn\"><span class=\"glyphicon glyphicon-export\" title=\"导出Excel\"></span>导出</span>")
	}
	if pagerbuttons:=param.ColConfigDict[len(param.ColConfigDict)-1];pagerbuttons.Tag=="pagerbuttons"{
		bodybuf.WriteString(pagerbuttons.Text)
	}
	bodybuf.WriteString("</div>")

	if style!=model.Style_Tree{
		bodybuf.WriteString("<div class=\"rt-pager-controls\">")
		bodybuf.WriteString("&nbsp;<span class=\"glyphicon glyphicon-step-backward rt-pager-firstPage\"></span>")
		bodybuf.WriteString("&nbsp;<span class=\"glyphicon glyphicon-backward rt-pager-prevPage\"></span>")
		bodybuf.WriteString("&nbsp;<span class=\"pager-separator\"></span>&nbsp;")
		bodybuf.WriteString("第&nbsp;<input type=\"text\" class=\"rt-pager-page\" value=\"")
		bodybuf.WriteString(function.IntToString(param.Settings.Page))
		bodybuf.WriteString("\"/>&nbsp;页，")
		bodybuf.WriteString("共&nbsp;<span class=\"rt-pager-totalPages\">")
		bodybuf.WriteString(function.IntToString(totalpages))
		bodybuf.WriteString("</span>&nbsp;页")
		bodybuf.WriteString("&nbsp;<span class=\"pager-separator\"></span>&nbsp;")
		bodybuf.WriteString("<span class=\"glyphicon glyphicon-forward rt-pager-nextPage\"></span>&nbsp;")
		bodybuf.WriteString("<span class=\"glyphicon glyphicon-step-forward rt-pager-lastPage\"></span>&nbsp;&nbsp;")
		bodybuf.WriteString("<select class=\"rt-pager-rowList\">")
		for _,v:=range rowlist{
			bodybuf.WriteString("<option value=\"")
			bodybuf.WriteString(v)
			bodybuf.WriteString("\"")
			if i,_:=function.StringToInt(v);i==param.Settings.Rows{
				bodybuf.WriteString(" selected")
			}
			bodybuf.WriteString(">")
			bodybuf.WriteString(v)
			bodybuf.WriteString("</option>")
		}
		bodybuf.WriteString("</select>")
		bodybuf.WriteString("</div>")
		bodybuf.WriteString("<div class=\"rt-pager-records\">第&nbsp;")
		bodybuf.WriteString(function.IntToString(start))
		bodybuf.WriteString(" - ")
		bodybuf.WriteString(function.IntToString(end))
		bodybuf.WriteString("&nbsp;条，")
		bodybuf.WriteString("共&nbsp;<span class=\"rt-pager-totalRecords\">")
		bodybuf.WriteString(function.IntToString(count))
		bodybuf.WriteString("</span>&nbsp;条</div>")
	}
	bodybuf.WriteString("</div>")
	return
}

func BuildSelectorBar(param *Param,size int,selectorbuf *bytes.Buffer,conditionbuf *bytes.Buffer)  (err error){
	for i:=0; i<size;i++  {
		selectordata:=make(map[string]string, 0)
		if !param.ColConfigDict[i].IsInselector{
			continue
		}
		being,selectordata:=db.HGetSelectorBarCache(param.TableConfig.Name,param.ColConfigDict[i].Tag)
		if being==true{
			fmt.Println(being,selectordata)
		}else {
			e:=db.SetSelectorBarCacheDuration(param.TableConfig.Name,param.ColConfigDict[i].Tag)
			if e!=nil{
				return  e
			}
			fmt.Println("nothing")
			query,_:=GetSelectQuery(param,"distinct("+param.ColConfigDict[i].Tag+")")
			result,e:=db.Query(query)
			if e!=nil{
				return e
			}
			for j:=0;result.Next();j++{
				var value string
				if e=result.Scan(&value);e!=nil{
					fmt.Println("BuildSelectorBar:",e.Error())
					return e
				}
				fmt.Println(value)
				selectordata[function.IntToString(j)]=value
				e:=db.HSetSelectorBarCache(param.TableConfig.Name,param.ColConfigDict[i].Tag,function.IntToString(j),value)
				if e!=nil{
					return e
				}
			}
		}
	}
	return  nil
}