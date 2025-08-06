package main

import (
	"context"
	"fmt"
	user "ex5_cleanarchi/kitex_gen/user"
	// "ex5_cleanarchi/internal/entity"
	"ex5_cleanarchi/internal/repository"
	"ex5_cleanarchi/internal/service"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, request *user.CreateUserRequest) (resp *user.User, err error) {
	// TODO: Your code here...
	newUser, err := service.CreateUser(request.Name, request.Email)
	if err != nil {
		return nil, err
	}
	if newUser == nil {
		return nil, fmt.Errorf("Create User: nil error")
	}

	return &user.User{
		Id: newUser.Id,
		Name: newUser.Name,
		Email: newUser.Email,
	}, nil
}

// Getuser implements the UserServiceImpl interface.
func (s *UserServiceImpl) Getuser(ctx context.Context, request *user.GetUserRequest) (resp *user.User, err error) {
	// TODO: Your code here...
	foundUser, err := repository.GetUser(request.Id)
	if err != nil {
		return nil, err
	}

	return &user.User{
		Id: foundUser.Id,
		Name: foundUser.Name,
		Email: foundUser.Email,
	}, nil
}
