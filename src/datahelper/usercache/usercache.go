package usercache
type UserDetail struct {
	Uid       uint32  `json:"uid"`
	Password  string  `json:"password"`
}
func GetUserDetail(uid uint32) (detail *UserDetail, e error) {
	return
}