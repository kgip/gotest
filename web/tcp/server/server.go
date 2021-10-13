package server

type Service interface {
	Registry(name string) (bool, error)
	GetName(id int64) (string, error)
}
