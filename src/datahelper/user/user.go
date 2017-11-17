package user

import (
	"fmt"
	"utils/function"
	"datahelper/usercache"
	"errors"
	"net/url"
	"strings"
	"datahelper/db"
)

type LoginSuccessData struct {
	Auth    string `json:"auth"`
	Avatar   string`json:"avatar"`
}

type LoginFailData struct {
	FailCode int    `json:"failCode"`
	Msg      string `json:"msg"`
}
//验证用户
func UserValid(uid uint32, hashcode string,useragent string) (valid bool, err error) {
	//验证
	ud,err :=usercache.GetUserDetail(uid)
	if err!=nil{
		return
	}
	//fmt.Println(ud.UserAgent+"||"+useragent)
	if strings.Index(useragent,ud.UserAgent)>-1{
		//fmt.Print("ture")
		token,e:=url.QueryUnescape(hashcode)
		if e != nil {
			return false,e
		}
		if token==function.Md5String(fmt.Sprintf("%s|%s", ud.Uid, ud.Password)){
			valid=true
		}else {
			err=errors.New("您的cookie失效了呢，请重新登录！")
		}
	}else {
		err=errors.New("您的登录IP不正确噢，怎么能偷盗人家帐号！")
	}
	return
}
func GetUidbyName(name string) (uid uint32, e error) {
	uid=0
	if(name=="wind"){
		uid=331805370
	}
	return
}
func CheckUserForbid(uid uint32) (forbid bool) {
	forbid,_= db.CheckUserForbid(uid)
	return
}
func CheckUserState(uid uint32) (state bool) {
	state,_=db.CheckUserState(uid)
	return
}
func CheckUserLoginErr(uid uint32) (forbid bool) {
	if cnt, _ := db.UserLoginErrCnt(uid); cnt >= 10 {
		db.SetUserForbid(uid)
		return true
	}else {
		return false
	}
}
func CheckAuth(uid uint32, password string) (ud *usercache.UserDetail, e error) {
	//fmt.Print(utils.Md5String(fmt.Sprintf("%s_%d",password,148360)))
	ud,e=usercache.GetUserDetail(uid)
	if e!=nil{
		return
	}
	if passwordm :=function.Md5String(fmt.Sprintf("%s_%d",password,ud.Salt));ud.Password==passwordm{
		return ud,nil
	}else {
		return nil,errors.New("您的密码貌似不正确哦！")
	}
}

func CreateSuccessResp(ud *usercache.UserDetail) (res map[string]interface{}) {
	res = make(map[string]interface{}, 0)
	res["loginstatus"] = 0
	var sdata LoginSuccessData
	sdata.Auth = fmt.Sprintf("%d_%s", ud.Uid, function.Md5String(fmt.Sprintf("%s|%s", ud.Uid, ud.Password)))
	sdata.Avatar=ud.Avatar
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