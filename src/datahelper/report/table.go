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