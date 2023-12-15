package repository

import "github.com/AlbertPuwadol/grpc-clean/pkg/entity"

type IRepository interface {
	Hello() string
	Tasks1() []entity.Entity
	Task2() string
	Task3() int
}

type repository struct{}

func NewRepository() *repository {
	return &repository{}
}

func (r repository) Hello() string {
	return "OK"
}

func (r repository) Tasks1() []entity.Entity {
	return []entity.Entity{{Start: 1, End: 2, Entity: "Hospital", Label: "PLACE"}, {Start: 3, End: 10, Entity: "Doctor", Label: "PERSON"}, {Start: 14, End: 20, Entity: "Nurse", Label: "PERSON"}}
}

func (r repository) Task2() string {
	return "POSITIVE"
}

func (r repository) Task3() int {
	return 100
}
