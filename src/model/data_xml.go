package model

type ColumnConfig struct {
	Tag             string//和数据库feild相同
	Text            string//显示的汉字
	HasBtn          bool
	HasDateformat   bool
	HasDefaultValue bool
	HasFormatter    bool
	HasFormatterR   bool
	//HasLinkTo bool
	HasNavName   bool
	HasPrecision bool
	//HasRegex bool
	HasSearchType    bool
	HasSelectorMulti bool
	HasTimeTransfer  bool
	HasSearchBtnIcon bool
	HasSearchBtnFunc bool
	HasSelectorFunc  bool
	HasSelectorText  bool
	//IsInPercentageform bool
	IsInSelector  bool
	Search4Admin  bool
	ISCheckBox    bool
	IsSearchAdv   bool
	HasPower      bool
	Power         uint8 //0为管理员
	SearchBtnIcon string//查询框图标
	SearchBtnFunc string//查询框图标的点击方法
	BtnIcon       string
	BtnFunc       string
	ColumnName    string
	DateFormat    string
	DefaultValue  string//若找不到数据 ，默认值显示
	Formatter     string//用$$::$$可代替自身value，如“我即是$$::$$”（其实可用FormatterR来实现）
	FormatterR    string//为difinition文件夹下formatter.go的方法名，通过反射调用,可自定义table元素显示内容
	//LinkTo string
	NavName      string
	TimeTransfer string
	Precision    string
	//RegexPattern string
	//RegexReplacement string
	SearchType       string
	Selector         string
	SelectorFunc     string//为difinition文件夹下difinition.go的方法名，通过反射调用，和SelectorFuncAgrs，结合使用，可自定义selector的显示
	SelectorFuncAgrs string//如"0,管理员,1,Geust"，表是如果数据库的value为0为显示为管理员，value为1，则显示为Geust
	SelectorText     string
	Visibility       string //暂时有效值table-none,table-hidden
	//Passedcol []string
}
type TableConfig struct {
	Name            string
	DefaultOrder    string
	HasDefaultOrder bool
	Excel           string //生成excel
	HasExcel        bool
	//AdminName string
	//HasAdminName bool
	Power    uint8
	HasPower bool
}
