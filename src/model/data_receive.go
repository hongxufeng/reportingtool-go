package model

type Settings struct {
	Uid uint32//通过cookie验证用户  得到的用户标识UID
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
	HR bool
}
