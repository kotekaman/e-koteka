package repository

type Repository interface {
	Save(data interface{}) error
}