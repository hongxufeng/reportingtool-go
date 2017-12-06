package definition

import (
	"errors"
	"strings"
)

type DefinitionData struct{}

func (*DefinitionData) GetDefinitionData(param string) (map[string]string, error) {
	definitionData := make(map[string]string)
	data := strings.Split(param, ",")
	if len(data)%2 != 0 {
		return definitionData, errors.New("selector-func-agrs must be even")
	}
	for i := 0; i < len(data); i += 2 {
		definitionData[data[i]] = data[i+1]
	}
	return definitionData, nil
}
