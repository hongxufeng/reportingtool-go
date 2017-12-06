package definition

import (
	"bytes"
	"model"
)

func (*Definition) PutInText(config model.ColumnConfig, cellValue string) (string, error) {
	var buf bytes.Buffer
	buf.WriteString("<input type=\"text\" class=\"rt-celltext\" value=\"")
	buf.WriteString(cellValue)
	buf.WriteString("\">")
	return buf.String(), nil
}
