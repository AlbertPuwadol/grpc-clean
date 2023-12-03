package usecase

import (
	"github.com/AlbertPuwadol/grpc-clean/pkg/entity"
	"github.com/AlbertPuwadol/grpc-clean/pkg/repository"
)

type IUsecase interface {
	Hello() string
	Tasks(string) entity.TaskResult
	Task1(string) entity.Task1
	Task2(string) string
	Task3(string) int
}

type usecase struct {
	repository repository.IRepository
}

func NewUsecase(repository repository.IRepository) *usecase {
	return &usecase{repository: repository}
}

func (u usecase) Hello() string {
	return u.repository.Hello()
}

func (u usecase) Tasks(text string) entity.TaskResult {
	return entity.TaskResult{
		Task1: entity.Task1{
			Text:     text,
			Entities: u.repository.Tasks1(),
		},
		Task2: u.repository.Task2(),
		Task3: u.repository.Task3(),
	}
}

func (u usecase) Task1(text string) entity.Task1 {
	return entity.Task1{
		Text:     text,
		Entities: u.repository.Tasks1(),
	}
}

func (u usecase) Task2(text string) string {
	return u.repository.Task2()
}

func (u usecase) Task3(text string) int {
	return u.repository.Task3()
}
