package model

import "time"

const (
	User_Forfid_Cnt=10//错误次数大于此值时，则判定为登陆频繁
	User_Forfid_Expiration_Duration=time.Minute * 10 //断定用户登陆频繁锁定的时间
	User_Info_Persistence_Duration=time.Hour * 14 //登录用户缓存持续时间
)