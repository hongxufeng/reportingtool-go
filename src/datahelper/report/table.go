package report

import (
	"bytes"
	"database/sql"
	"strings"
)

func GetTableSearch(param *Param) string {
	var buf bytes.Buffer
	buf.WriteString("")
	return buf.String()
}

func BuildTableHead(param *Param, rows *sql.Rows) string {
	var buf bytes.Buffer
	buf.WriteString("<thead>")
	buf.WriteString("<tr>")
	if param.Settings.HasCheckbox {
		buf.WriteString("<th class=\\\"rt-th-checkbox\\\" name=\\\"rt-th-checkbox\\\">")
		buf.WriteString("<div class=\\\"rt-checkboxWrapper\\\">")
		buf.WriteString("<input type=\\\"checkbox\\\" class=\\\"rt-checkbox\\\"/>")
		buf.WriteString("</div>")
		buf.WriteString("</th>")
	}
	columns, _ := rows.Columns()
	for i, _ := range columns {
		if param.ColConfigDict[i].Visibility == "none" {
			continue;
		}
		buf.WriteString("<th ")
		buf.WriteString("class=\\\"")
		if param.ColConfigDict[i].Visibility == "hidden" {
			buf.WriteString("hiddenCol");
		} else {
			buf.WriteString("rt-sort");
		}
		buf.WriteString("\\\"")
		buf.WriteString(" name=\\\"")
		buf.WriteString(param.ColConfigDict[i].Tag)
		buf.WriteString("\\\">")
		buf.WriteString(param.ColConfigDict[i].Text)
		order := strings.Split(param.Settings.Order, " ")
		if len(order) > 1 && order[0] == param.ColConfigDict[i].Tag {
			buf.WriteString("<span class=\\\"glyphicon glyphicon-")
			if strings.ToLower(order[1]) == "asc" {
				buf.WriteString("arrow-up")
			} else {
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

func BuildTableBody(param *Param, rows *sql.Rows) string {
	var buf bytes.Buffer
	var checkvalue string
	buf.WriteString("<tbody>")

	columns, _ := rows.Columns()
	size := len(columns)

	for i := 0; i < size; i++ {
		if param.ColConfigDict[i].ISCheckBox {
			checkvalue = param.ColConfigDict[i].Text
			break
		}
	}
	for i := 0; rows.Next(); i++ {
		buf.WriteString("<tr>")
		if param.Settings.HasCheckbox {
			buf.WriteString("<td class=\\\"rt-td-checkbox\\\" name=\\\"rt-td-checkbox\\\" data-value=\\\"")
			buf.WriteString(checkvalue)
			buf.WriteString("\\\">")
			buf.WriteString("<input type=\\\"checkbox\\\"  class=\\\"rt-checkbox\\\" value=\\\"")
			buf.WriteString(checkvalue)
			buf.WriteString("\\\" />")
			buf.WriteString("</div>")
			buf.WriteString("</td>")
		}
		var s [size]string
		rows.Scan(s...)
		for i := 0; i < size; i++ {
			if param.ColConfigDict[i].Visibility == "none" {
				continue
			}
			buf.WriteString("<td name=\\\"")
			buf.WriteString(param.ColConfigDict[i].Tag)
			buf.WriteString("\\\">")
			if param.ColConfigDict[i].Visibility == "hidden" {
				buf.WriteString(" class=\\\"hiddenCol\\\"")
			}
			cell:=s[i]
			buf.WriteString(" data-value=\\\"")
			buf.WriteString(cell)
			buf.WriteString("\\\">")
			if param.ColConfigDict[i].HasFormatter {
				//反射查找相应函数
				//cellValue = FormatCell(dataTable.Columns, row, colConfig.Formatter, colName, cellValue);
			}
			buf.WriteString(cell)
			if param.ColConfigDict[i].HasBtn {
				buf.WriteString("<span class=\\\"rt-cell-btn glyphicon glyphicon-")
				buf.WriteString(param.ColConfigDict[i].BtnIcon)
				buf.WriteString("\\\" onclick=\\\"")
				buf.WriteString(param.ColConfigDict[i].BtnFunc)
				buf.WriteString("\\\"></span>")
			}
			buf.WriteString("</td>")
		}
		buf.WriteString("</tr>")
	}
	buf.WriteString("</tbody>")
	return buf.String()
}

func GetTableBody(param *Param, rows *sql.Rows) string {
	var buf bytes.Buffer
	buf.WriteString("<table class=\\\"table table-condensed\\\">")
	buf.WriteString(BuildTableHead(param, rows))
	buf.WriteString(BuildTableBody(param, rows))
	return buf.String()
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
