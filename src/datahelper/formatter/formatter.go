package formatter

import (
	"bytes"
	"model"
)

type CellFormatter struct{}

func (*CellFormatter) PutInText(config model.ColumnConfig, cellValue string) (string, error) {
	var buf bytes.Buffer
	buf.WriteString("<input type=\"text\" class=\"rt-celltext\" value=\"")
	buf.WriteString(cellValue)
	buf.WriteString("\">")
	return buf.String(), nil
}
