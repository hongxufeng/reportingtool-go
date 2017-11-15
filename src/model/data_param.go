package model

const (
	CMD_GetTable="GetTable"
	CMD_SearchTree="SearchTree"
	CMD_LocateNode="LocateNode"
	Style_Rect="rect"
	Style_Table = "table"
	Style_Tree = "tree"
)

type Param struct {
	XmlTable interface{}//XML获得的数据结构
	Settings Settings
	Uid uint32
	IsAdmin bool  //用户判断得到的权限
	ColConfigDict []ColumnConfig
}


