package model

type Settings struct {
	Cmd string
	Style string
	TableID string
	ConfigFile string
	HasCheckbox string
	RowList string
	Condition string
	Page int
	Rows int
	ColPage int
}

type LoginData struct {
	Username string
	Password string
}