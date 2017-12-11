package report

import (
	"bytes"
	"database/sql"
	"datahelper/definition"
	"errors"
	"fmt"
	"model"
	"reflect"
	"strings"
	"utils/function"
	"utils/service"
)

func GetTable(req *service.HttpRequest, param *Param, settings *model.Settings, rows *sql.Rows, size int, bodybuf *bytes.Buffer, searchbuf *bytes.Buffer, rowbuf *bytes.Buffer, count int) (err error) {
	bodybuf.WriteString("<table class=\"table table-condensed\">")
	err = BuildTableHead(req, param, settings, size, bodybuf, searchbuf)
	if err != nil {
		return
	}
	err = BuildTableBody(param, settings, rows, size, bodybuf)
	if err != nil {
		return
	}

	bodybuf.WriteString("</table>")

	err = BuildTablePager(param, settings, bodybuf, count, model.Style_Table)
	if err != nil {
		return
	}
	err = BuildNullRow(param, settings, size, rowbuf)
	if err != nil {
		return
	}
	return
}

func BuildSearchingBlock(req *service.HttpRequest, columnconfig *model.ColumnConfig, searchbuf *bytes.Buffer) (err error) {
	if !columnconfig.HasSearchType {
		return
	}
	searchbuf.WriteString("<div")
	if columnconfig.IsSearchAdv {
		searchbuf.WriteString(" class=\"rt-search-adv\"")
	}
	searchbuf.WriteString("><span class=\"rt-search-heading\">")
	searchbuf.WriteString(columnconfig.Text)
	searchbuf.WriteString("：</span>")
	if columnconfig.SearchType == "range" || columnconfig.SearchType == "date" {
		var start, end string
		_ = req.GetParams(false, columnconfig.Tag+">=", &start, columnconfig.Tag+"<+", &end)
		searchbuf.WriteString("<input type=\"text\" class=\"rt-search-txt form-control ")
		searchbuf.WriteString(columnconfig.SearchType)
		searchbuf.WriteString("\" name=\"")
		searchbuf.WriteString(columnconfig.Tag)
		searchbuf.WriteString("\" data-sign=\"%3e%3d\" value=\"")
		searchbuf.WriteString(start)
		searchbuf.WriteString("\"/>")
		searchbuf.WriteString("<span class=\"search-span-minus\"> - </span>")
		searchbuf.WriteString("<input type=\"text\" class=\"rt-search-txt form-control ")
		searchbuf.WriteString(columnconfig.SearchType)
		searchbuf.WriteString("\" name=\"")
		searchbuf.WriteString(columnconfig.Tag)
		searchbuf.WriteString("\" data-sign=\"%3e%3d\" value=\"")
		searchbuf.WriteString(end)
		searchbuf.WriteString("\"/>")
	} else {
		searchbuf.WriteString("<input type=\"text\" class=\"rt-search-txt form-control")
		var value string
		_ = req.GetParams(false, columnconfig.Tag, &value)
		if len(value) == 0 {
			_ = req.GetParams(false, columnconfig.Tag+"~~", &value)
		}
		searchbuf.WriteString("\" name=\"")
		searchbuf.WriteString(columnconfig.Tag)
		searchbuf.WriteString("\" data-sign=\"%7e%7e\" value=\"")
		searchbuf.WriteString(value)
		searchbuf.WriteString("\"/>")
		if columnconfig.HasSearchBtnFunc && columnconfig.HasSearchBtnIcon {
			searchbuf.WriteString("<span class=\"glyphicon glyphicon-")
			searchbuf.WriteString(columnconfig.SearchBtnIcon)
			searchbuf.WriteString(" rt-search-txt-btn\" onclick=\"")
			searchbuf.WriteString(columnconfig.SearchBtnFunc)
			searchbuf.WriteString("\"></span>")
		}
	}
	searchbuf.WriteString("</div>")
	return
}

func BuildTableHead(req *service.HttpRequest, param *Param, settings *model.Settings, size int, bodybuf *bytes.Buffer, searchbuf *bytes.Buffer) (err error) {
	bodybuf.WriteString("<thead>")
	bodybuf.WriteString("<tr>")
	if settings.HasCheckbox {
		bodybuf.WriteString("<th class=\"rt-th-checkbox\" name=\"rt-th-checkbox\">")
		bodybuf.WriteString("<div class=\"rt-checkboxWrapper\">")
		bodybuf.WriteString("<input type=\"checkbox\" class=\"rt-checkbox\"/>")
		bodybuf.WriteString("</div>")
		bodybuf.WriteString("</th>")
	}
	fmt.Println("size:", size)
	for i := 0; i < size; i++ {
		if strings.Index(param.ColConfigDict[i].Visibility, "table-none") > -1 {
			continue
		}
		err = BuildSearchingBlock(req, &param.ColConfigDict[i], searchbuf)
		//fmt.Println("searchbuf:",searchbuf.String())
		if err != nil {
			return
		}
		bodybuf.WriteString("<th ")
		bodybuf.WriteString("class=\"")
		if strings.Index(param.ColConfigDict[i].Visibility, "table-hidden") > -1 {
			bodybuf.WriteString("hiddenCol")
		} else {
			bodybuf.WriteString("rt-sort")
		}
		bodybuf.WriteString("\"")
		bodybuf.WriteString(" name=\"")
		bodybuf.WriteString(param.ColConfigDict[i].Tag)
		bodybuf.WriteString("\">")
		bodybuf.WriteString(param.ColConfigDict[i].Text)
		order := strings.Split(settings.Order, " ")
		if len(order) > 1 && order[0] == param.ColConfigDict[i].Tag {
			bodybuf.WriteString("<span class=\"glyphicon glyphicon-")
			if strings.ToLower(order[1]) == "asc" {
				bodybuf.WriteString("arrow-up")
			} else {
				bodybuf.WriteString("arrow-down")
			}
			bodybuf.WriteString("\"></span>")
		}
		bodybuf.WriteString("</th>")
	}
	for i := size; i < size+2; i++ {
		if param.ColConfigDict[i].Tag == "buttons" {
			bodybuf.WriteString("<th name=\"操作\">")
			bodybuf.WriteString("操作")
			bodybuf.WriteString("</th>")
		}
	}
	bodybuf.WriteString("</tr>")
	bodybuf.WriteString("</thead>")
	fmt.Println(bodybuf.String())
	return
}

func BuildTableBody(param *Param, settings *model.Settings, rows *sql.Rows, size int, bodybuf *bytes.Buffer) (err error) {
	var checkvalue string
	bodybuf.WriteString("<tbody>")

	var s []interface{}
	for i := 0; i < size; i++ {
		var white = ""
		var p *string
		p = &white
		s = append(s, p)
	}

	for i := 0; i < size; i++ {
		if param.ColConfigDict[i].ISCheckBox {
			checkvalue = param.ColConfigDict[i].Text
			break
		}
	}
	for rows.Next() {
		bodybuf.WriteString("<tr>")
		if settings.HasCheckbox {
			bodybuf.WriteString("<td class=\"rt-td-checkbox\" name=\"rt-td-checkbox\" data-value=\"")
			bodybuf.WriteString(checkvalue)
			bodybuf.WriteString("\">")
			bodybuf.WriteString("<input type=\"checkbox\"  class=\"rt-checkbox\" value=\"")
			bodybuf.WriteString(checkvalue)
			bodybuf.WriteString("\" />")
			bodybuf.WriteString("</div>")
			bodybuf.WriteString("</td>")
		}
		fmt.Println(s...)
		err = rows.Scan(s...)
		if err != nil {
			return
		}
		fmt.Println(s)
		fmt.Println(function.PArrayToSArray(s))
		for i := 0; i < size; i++ {
			if strings.Index(param.ColConfigDict[i].Visibility, "table-none") > -1 {
				continue
			}
			if param.ColConfigDict[i].HasPower && param.Power >= param.ColConfigDict[i].Power {
				continue
			}
			bodybuf.WriteString("<td name=\"")
			bodybuf.WriteString(param.ColConfigDict[i].Tag)
			bodybuf.WriteString("\"")
			if strings.Index(param.ColConfigDict[i].Visibility, "table-hidden") > -1 {
				bodybuf.WriteString(" class=\"hiddenCol\"")
			}
			cell := function.ToString(s[i])
			fmt.Println(cell)
			bodybuf.WriteString(" data-value=\"")
			bodybuf.WriteString(cell)
			bodybuf.WriteString("\">")
			cell, _ = Format(&param.ColConfigDict[i], cell)
			bodybuf.WriteString(cell)
			if param.ColConfigDict[i].HasBtn {
				bodybuf.WriteString("<span class=\"rt-cell-btn glyphicon glyphicon-")
				bodybuf.WriteString(param.ColConfigDict[i].BtnIcon)
				bodybuf.WriteString("\" onclick=\"")
				bodybuf.WriteString(param.ColConfigDict[i].BtnFunc)
				bodybuf.WriteString("\"></span>")
			}
			bodybuf.WriteString("</td>")
			if i == size-1 {
				if param.ColConfigDict[i+1].Tag == "buttons" {
					bodybuf.WriteString("<th name=\"操作\">")
					bodybuf.WriteString(param.ColConfigDict[i+1].Text)
					bodybuf.WriteString("</th>")
				} else if param.ColConfigDict[i+2].Tag == "buttons" {
					bodybuf.WriteString("<th name=\"操作\">")
					bodybuf.WriteString(param.ColConfigDict[i+2].Text)
					bodybuf.WriteString("</th>")
				}
			}
		}
		bodybuf.WriteString("</tr>")
	}
	bodybuf.WriteString("</tbody>")
	return
}

func BuildNullRow(param *Param, settings *model.Settings, size int, rowbuf *bytes.Buffer) (err error) {
	var definitionData definition.Definition
	define := reflect.ValueOf(&definitionData)
	rowbuf.WriteString("<tr>")
	if settings.HasCheckbox {
		rowbuf.WriteString("<td class=\"rt-td-checkbox\" name=\"rt-td-checkbox\" data-value=\"\">")
		rowbuf.WriteString("<div class=\"rt-checkboxWrapper\">")
		rowbuf.WriteString("<input type=\"checkbox\"  class=\"rt-checkbox\" value=\"\" />")
		rowbuf.WriteString("</div>")
		rowbuf.WriteString("</td>")
	}
	for i := 0; i < size; i++ {
		if strings.Index(param.ColConfigDict[i].Visibility, "table-none") > -1 {
			continue
		}
		rowbuf.WriteString("<td name=\"")
		rowbuf.WriteString(param.ColConfigDict[i].Tag)
		rowbuf.WriteString("\"")
		if strings.Index(param.ColConfigDict[i].Visibility, "table-hidden") > -1 {
			rowbuf.WriteString(" class=\"hiddenCol\"")
		}
		cell := function.ToString(param.ColConfigDict[i].Text)
		if param.ColConfigDict[i].HasDefaultValue {
			cell = param.ColConfigDict[i].DefaultValue
		}
		fmt.Println(cell)
		rowbuf.WriteString(" data-value=\"")
		rowbuf.WriteString(cell)
		rowbuf.WriteString("\">")
		if param.ColConfigDict[i].HasFormatter {
			//反射查找相应函数
			method := define.MethodByName(param.ColConfigDict[i].Formatter)
			values := method.Call([]reflect.Value{reflect.ValueOf(param.ColConfigDict[i]), reflect.ValueOf(cell)})
			fmt.Println(values)
			if len(values) != 2 {
				err = errors.New(fmt.Sprintf("method %s return value is not 2.", param.ColConfigDict[i].Formatter))
			}
			switch x := values[1].Interface().(type) {
			case nil:
				cell = values[0].String()
			default:
				return x.(error)
			}
			//cellValue = FormatCell(dataTable.Columns, row, colConfig.Formatter, colName, cellValue);
		}
		rowbuf.WriteString(cell)
		if param.ColConfigDict[i].HasBtn {
			rowbuf.WriteString("<span class=\"rt-cell-btn glyphicon glyphicon-")
			rowbuf.WriteString(param.ColConfigDict[i].BtnIcon)
			rowbuf.WriteString("\" onclick=\"")
			rowbuf.WriteString(param.ColConfigDict[i].BtnFunc)
			rowbuf.WriteString("\"></span>")
		}
		rowbuf.WriteString("</td>")
		if i == size-1 {
			if param.ColConfigDict[i+1].Tag == "buttons" {
				rowbuf.WriteString("<th name=\"操作\">")
				rowbuf.WriteString(param.ColConfigDict[i+1].Text)
				rowbuf.WriteString("</th>")
			} else if param.ColConfigDict[i+2].Tag == "buttons" {
				rowbuf.WriteString("<th name=\"操作\">")
				rowbuf.WriteString(param.ColConfigDict[i+2].Text)
				rowbuf.WriteString("</th>")
			}
		}
	}
	rowbuf.WriteString("</tr>")

	return
}
