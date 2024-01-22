package store

type OrderRepository interface {
	Create() error
	FindById() error
}
