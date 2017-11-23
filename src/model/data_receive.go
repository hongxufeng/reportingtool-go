package model

type Settings struct {
	Style string
	TableID string
	ConfigFile string
	HasCheckbox bool
	RowList string
	Condition string
	Page int
	Rows int
	ColPage int
	Order string
}

type LoginData struct {
	Username string
	Password string
}