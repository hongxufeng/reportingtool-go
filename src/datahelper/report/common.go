package report

import (
	"bytes"
	"model"
	"utils/function"
	"strings"
)

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