package service

type Service interface {
	Start()
	Stop()
	Status() string
}
