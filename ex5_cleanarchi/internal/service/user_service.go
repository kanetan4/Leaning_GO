package service

import (
	"ex5_cleanarchi/internal/entity"
	"ex5_cleanarchi/internal/repository"
)

func CreateUser(name string, email string) (*entity.User, error) {
	return repository.SaveUser(name, email)
}

func GetUser(id int64) (*entity.User, error) {
	return repository.GetUser(id)
}