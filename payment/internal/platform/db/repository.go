package db

type Repository[T any, ID comparable] interface {
	GetAll() ([]T, error)
	GetByID(id ID) (T, error)
	Create(entity T) (T, error)
	Update(id ID, entity T) (T, error)
	Delete(id ID) error
}