package service

import "utils/config"

type Module interface {
	Init(config *config.Config) (err error)
}