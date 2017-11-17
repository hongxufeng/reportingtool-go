package service

import "utils/config"

type LEVEL byte

const (
	PROD LEVEL = iota//正式环境
	PRE//预发布环境
	TEST//测试环境
	DEV//开发环境
	OFF
)
type Module interface {
	Init(config *config.Config) (err error)
}