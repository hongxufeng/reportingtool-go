package model

type Settings struct {
	CMD string`json:"cmd"`
	Style string`json:"style"`
	TableID string`json:"table"`
	ConfigFile string`json:"configFile"`
	HasCheckbox string `json:"hasCheckbox"`
	RowList string`json:"rowList"`
	Condition string`json:"condition"`
	QueryString QueryString`json:"queryString"`
}
type QueryString struct {

}