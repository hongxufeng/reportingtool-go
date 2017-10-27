package model

type Settings struct {
	CMD string`json:"cmd"`
	Style string`json:"style"`
	TableID string`json:"table"`
	ConfigFile string`json:"configFile"`
	HasCheckbox string `json:"hasCheckbox"`
	RowList string`json:"rowList"`
	Condition string`json:"condition"`
	Page int `json:"page"`
	Rows int`json:"rows"`
	ColPage int`json:"colpage"`
	HR bool `json:"hr"`
}
