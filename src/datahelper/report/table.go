package report

import (
	"bytes"
	"database/sql"
)

func GetTableSearch(param *Param) string{
	var buf bytes.Buffer
	buf.WriteString("")
	return buf.String()
}
func GetTableBody(param *Param,rows *sql.Rows) string{
	var buf bytes.Buffer
	buf.WriteString("<table class=\\\"table table-condensed\\\">")
	buf.WriteString("<thead>")
	buf.WriteString("<tr>")
	if param.Settings.HasCheckbox{
		buf.WriteString("<th class=\\\"rt-th-checkbox\\\" name=\\\"rt-th-checkbox\\\">")
		buf.WriteString("<div class=\\\"rt-checkboxWrapper\\\">")
		buf.WriteString("<input type=\\\"checkbox\\\" class=\\\"rt-checkbox\\\"/>")
		buf.WriteString("</div>")
		buf.WriteString("</th>")
	}

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