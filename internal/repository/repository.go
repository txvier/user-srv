package repository

type Repository interface {
	Get(id string) interface{}
}
