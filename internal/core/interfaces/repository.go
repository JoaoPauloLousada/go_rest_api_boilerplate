package interfaces

type Repository[T any] interface {
	GetAll() ([]T, error)
	Create(entity T) error
}
