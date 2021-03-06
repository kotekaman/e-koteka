package repository

type Repository interface {
	Save(data interface{}) error
	GetAll(params ...interface{}) (interface{}, error)
	Get(params ...interface{}) (interface{}, error)
	Delete(params ...interface{}) error
	Update(params ...interface{}) error
}
