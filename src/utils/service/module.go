package service

type Module interface {
	Init() (err error)
}
