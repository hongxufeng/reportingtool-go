package user

import (
	"fmt"
	"utils"
	"datahelper/usercache"
)
type LoginSuccessData struct {
	Uid      uint32 `json:"uid"`
	Token    string `json:"token"`
}

type LoginFailData struct {
	FailCode int    `json:"failCode"`
	Msg      string `json:"msg"`
}
//验证用户
func UserValid(uid uint32, hashcode string) (valid bool, e error) {
	//验证
	valid=true
	return
}
func CheckUserForbid(username string) (valid bool, e error) {
	//验证
	valid=true
	return
}
func CheckUserState(username string) (valid bool, e error) {
	//验证
	valid=true
	return
}
func CheckAuth(name string, password string) (ud *usercache.UserDetail, e error) {
	return
}

func CreateSuccessResp(ud *usercache.UserDetail) (res map[string]interface{}, e error) {
	res = make(map[string]interface{}, 0)
	res["loginstatus"] = 0
	var sdata LoginSuccessData
	sdata.Token = fmt.Sprintf("%d_%s", ud.Uid, utils.Md5String(fmt.Sprintf("%s|%s", ud.Uid, ud.Password)))
	res["userdata"] = sdata
	return
}

func CreateFailResp(code int, msg string) (res map[string]interface{}) {
	res = make(map[string]interface{}, 0)
	var fdata LoginFailData
	res["loginstatus"] = 1
	fdata.FailCode = code
	fdata.Msg = msg
	res["faildata"] = fdata
	return
}