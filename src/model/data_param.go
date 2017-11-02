package model

type Param struct {
	XmlTable interface{}//XML获得的数据结构
	Settings Settings
	Uid uint32
	IsAdmin bool  //用户判断得到的权限
	ColConfigDict []ColumnConfig
}

const (
	Rect="rect"
	Table = "table"
	Tree = "tree"
)

