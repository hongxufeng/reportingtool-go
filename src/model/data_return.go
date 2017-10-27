package model

type TableModel struct {
	Search string`json:"search"`
	Body string`json:"body"`
	Selector string`json:"selector"`
	Condition string`json:"condition"`
	Row string `json:"row"`
}