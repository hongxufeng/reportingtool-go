package report

import (
	"bytes"
	"database/sql"
	"strings"
	"utils/function"
	"fmt"
)

func GetTableSearch(param *Param) string {
	var buf bytes.Buffer
	buf.WriteString("")
	return buf.String()
}

func BuildTableHead(param *Param, rows *sql.Rows,bodybuf *bytes.Buffer) (err error) {
	bodybuf.WriteString("<thead>")
	bodybuf.WriteString("<tr>")
	if param.Settings.HasCheckbox {
		bodybuf.WriteString("<th class=\\\"rt-th-checkbox\\\" name=\\\"rt-th-checkbox\\\">")
		bodybuf.WriteString("<div class=\\\"rt-checkboxWrapper\\\">")
		bodybuf.WriteString("<input type=\\\"checkbox\\\" class=\\\"rt-checkbox\\\"/>")
		bodybuf.WriteString("</div>")
		bodybuf.WriteString("</th>")
	}
	columns, _ := rows.Columns()
	size:=len(columns)
	fmt.Println(size)
	for i:=0;i<size;i++ {
		if param.ColConfigDict[i].Visibility == "none" {
			continue;
		}
		bodybuf.WriteString("<th ")
		bodybuf.WriteString("class=\\\"")
		if param.ColConfigDict[i].Visibility == "hidden" {
			bodybuf.WriteString("hiddenCol");
		} else {
			bodybuf.WriteString("rt-sort");
		}
		bodybuf.WriteString("\\\"")
		bodybuf.WriteString(" name=\\\"")
		bodybuf.WriteString(param.ColConfigDict[i].Tag)
		bodybuf.WriteString("\\\">")
		bodybuf.WriteString(param.ColConfigDict[i].Text)
		order := strings.Split(param.Settings.Order, " ")
		if len(order) > 1 && order[0] == param.ColConfigDict[i].Tag {
			bodybuf.WriteString("<span class=\\\"glyphicon glyphicon-")
			if strings.ToLower(order[1]) == "asc" {
				bodybuf.WriteString("arrow-up")
			} else {
				bodybuf.WriteString("arrow-down")
			}
			bodybuf.WriteString("\\\"></span>")
		}
		bodybuf.WriteString("</th>")
	}
	for i:=size;i<size+2;i++{
		if param.ColConfigDict[i].Tag=="buttons" {
			bodybuf.WriteString("<th name=\\\"操作\\\">")
			bodybuf.WriteString(param.ColConfigDict[i].Text)
			bodybuf.WriteString("</th>")
		}
	}
	bodybuf.WriteString("</tr>")
	bodybuf.WriteString("</thead>")
	fmt.Println(bodybuf.String())
	return
}

func BuildTableBody(param *Param, rows *sql.Rows,bodybuf *bytes.Buffer) (err error) {
	var checkvalue string
	bodybuf.WriteString("<tbody>")

	columns, _ := rows.Columns()
	size := len(columns)
	var s []interface{}
	for i:=0;i<size;i++ {
		var white =""
		var p *string
		p=&white
		s=append(s,p)
	}

	for i := 0; i < size; i++ {
		if param.ColConfigDict[i].ISCheckBox {
			checkvalue = param.ColConfigDict[i].Text
			break
		}
	}
	for i := 0; rows.Next(); i++ {
		bodybuf.WriteString("<tr>")
		if param.Settings.HasCheckbox {
			bodybuf.WriteString("<td class=\\\"rt-td-checkbox\\\" name=\\\"rt-td-checkbox\\\" data-value=\\\"")
			bodybuf.WriteString(checkvalue)
			bodybuf.WriteString("\\\">")
			bodybuf.WriteString("<input type=\\\"checkbox\\\"  class=\\\"rt-checkbox\\\" value=\\\"")
			bodybuf.WriteString(checkvalue)
			bodybuf.WriteString("\\\" />")
			bodybuf.WriteString("</div>")
			bodybuf.WriteString("</td>")
		}
		fmt.Println(s...)
		rows.Scan(s...)
		fmt.Println(s)
		for i := 0; i < size; i++ {
			if param.ColConfigDict[i].Visibility == "none" {
				continue
			}
			bodybuf.WriteString("<td name=\\\"")
			bodybuf.WriteString(param.ColConfigDict[i].Tag)
			bodybuf.WriteString("\\\">")
			if param.ColConfigDict[i].Visibility == "hidden" {
				bodybuf.WriteString(" class=\\\"hiddenCol\\\"")
			}
			cell:=function.ToString(s[i])
			fmt.Println(cell)
			bodybuf.WriteString(" data-value=\\\"")
			bodybuf.WriteString(cell)
			bodybuf.WriteString("\\\">")
			if param.ColConfigDict[i].HasFormatter {
				//反射查找相应函数
				//cellValue = FormatCell(dataTable.Columns, row, colConfig.Formatter, colName, cellValue);
			}
			bodybuf.WriteString(cell)
			if param.ColConfigDict[i].HasBtn {
				bodybuf.WriteString("<span class=\\\"rt-cell-btn glyphicon glyphicon-")
				bodybuf.WriteString(param.ColConfigDict[i].BtnIcon)
				bodybuf.WriteString("\\\" onclick=\\\"")
				bodybuf.WriteString(param.ColConfigDict[i].BtnFunc)
				bodybuf.WriteString("\\\"></span>")
			}
			bodybuf.WriteString("</td>")
		}
		bodybuf.WriteString("</tr>")
	}
	for i:=size;i<size+2;i++{
		if param.ColConfigDict[i].Tag=="buttons" {
			bodybuf.WriteString("<th name=\\\"操作\\\">")
			bodybuf.WriteString("操作")
			bodybuf.WriteString("</th>")
		}
	}
	bodybuf.WriteString("</tbody>")
	return
}

func GetTable(param *Param, rows *sql.Rows, bodybuf *bytes.Buffer) (err error){
	bodybuf.WriteString("<table class=\\\"table table-condensed\\\">")
	err=BuildTableHead(param, rows,bodybuf)
	if err != nil {
		return
	}
	err=BuildTableBody(param, rows,bodybuf)
	if err != nil {
		return
	}
	bodybuf.WriteString("</table>")
	return
}

func GetTableSelector(param *Param) string {
	var buf bytes.Buffer
	return buf.String()
}
func GetTableCondition(param *Param) string {
	var buf bytes.Buffer
	return buf.String()
}
func GetTableRow(param *Param) string {
	var buf bytes.Buffer
	return buf.String()
}
