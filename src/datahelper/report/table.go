package report

import (
	"bytes"
	"database/sql"
	"fmt"
	"strings"
)

func GetTableSearch(param *Param) string{
	var buf bytes.Buffer
	buf.WriteString("")
	return buf.String()
}
func BuildTableHead(param *Param,rows *sql.Rows) string{
	var buf bytes.Buffer
	buf.WriteString("<thead>")
	buf.WriteString("<tr>")
	if param.Settings.HasCheckbox{
		buf.WriteString("<th class=\\\"rt-th-checkbox\\\" name=\\\"rt-th-checkbox\\\">")
		buf.WriteString("<div class=\\\"rt-checkboxWrapper\\\">")
		buf.WriteString("<input type=\\\"checkbox\\\" class=\\\"rt-checkbox\\\"/>")
		buf.WriteString("</div>")
		buf.WriteString("</th>")
	}
	columns,_:=rows.Columns()
	size:=len(columns)
	for i:=0;rows.Next();i++{
		if param.ColConfigDict[i].Visibility == "none" {
			continue;
		}
		var s [size]string
		rows.Scan(s...)
		buf.WriteString("<th ")
		if param.ColConfigDict[i].Tag!="buttons"||param.ColConfigDict[i].Tag!="pagerbuttons"{
			buf.WriteString("class=\\\"")
			if param.ColConfigDict[i].Visibility == "hidden" {
				buf.WriteString("hiddenCol");
			} else {
				buf.WriteString("rt-sort");
			}
			buf.WriteString("\\\"")
		}
		buf.WriteString(" name=\\\"")
		buf.WriteString(param.ColConfigDict[i].Tag)
		buf.WriteString("\\\">")
		buf.WriteString(param.ColConfigDict[i].Text)
		order:=strings.Split(param.Settings.Order," ")
		if len(order)>1&&order[0]==param.ColConfigDict[i].Tag{
			buf.WriteString("<span class=\\\"glyphicon glyphicon-")
			if strings.ToLower(order[1])=="asc" {
				buf.WriteString("arrow-up")
			}else {
				buf.WriteString("arrow-down")
			}
			buf.WriteString("\\\"></span>")
		}
		buf.WriteString("</th>")
	}
	buf.WriteString("</tr>")
	buf.WriteString("</thead>")
	return buf.String()
}
func GetTableBody(param *Param,rows *sql.Rows) string{
	var buf bytes.Buffer
	buf.WriteString("<table class=\\\"table table-condensed\\\">")
	buf.WriteString(BuildTableHead(param,rows))
	return buf.String()
}
func GetTableSelector(param *Param) string{
	var buf bytes.Buffer
	return buf.String()
}
func GetTableCondition(param *Param) string{
	var buf bytes.Buffer
	return buf.String()
}
func GetTableRow(param *Param) string{
	var buf bytes.Buffer
	return buf.String()
}