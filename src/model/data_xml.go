package model

type ColumnConfig struct {
	Tag string
	Text string
	HasBtn bool
	HasDateformat bool
	HasDefaultValue bool
	HasFormatter bool
	HasFormatterR bool
	//HasLinkTo bool
	HasNavname bool
	HasPrecision bool
	HasRegex bool
	HasSearchType bool
	HasSelectorMulti bool
	HasTimetransfer bool
	HasSearchBtnIcon bool
	HasSearchBtnFunc bool
	IsInPercentageform bool
	IsInselector bool
	Search4Admin bool
	ISCheckBox bool
	ISSearchAdv bool
	SearchBtnIcon string
	SearchBtnFunc string
	BtnIcon string
	BtnFunc string
	ColumnName string
	Dateformat string
	DefaultValue string
	Formatter string
	FormatterR string
	//LinkTo string
	Navname string
	Timetransfer string
	Precision string
	RegexPattern string
	RegexReplacement string
	SearchType string
	Selector string
	SelectorFunc string
	SelectorText string
	Visibility string
	Passedcol []string
}
type TableConfig struct {
	Name string
	DefaultOrder string
	HasDefaultOrder bool
	Excel string//生成excel
	HasExcel bool
	AdminName string
	HasAdminName bool
	Power uint8
	HasPower bool
}