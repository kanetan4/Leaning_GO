package repository

import (
	"ex5_cleanarchi/internal/entity"
	"encoding/json"
	"os"
	"sync"
	"fmt"
)

const fileName = "store.json"

var mutex sync.Mutex

func SaveUser(name string, email string) (*entity.User, error) {
	mutex.Lock()
	defer mutex.Unlock()

	users, err := LoadUsers()
	if err != nil {
		return nil, fmt.Errorf("LoadUsers Failed")
	}

	newID := int64(1)
	for _, user := range users {
		if user.Id >= newID {
			newID = user.Id + 1
		}
	}

	newUser := &entity.User{
		Id: newID,
		Name: name,
		Email: email,
	}

	users = append(users, newUser)

	err = SaveUsers(users)
	if err != nil {
		return nil, fmt.Errorf("SaveUsers Failed")
	}
	return newUser, nil
}

func LoadUsers() ([]*entity.User, error) {
	var users []*entity.User
	data, err := os.ReadFile(fileName)

	if err != nil {
		if os.IsNotExist(err) {
			return []*entity.User{}, nil
		}
		return nil, err
	}

	if len(data) == 0 {
		return []*entity.User{}, nil
	}

	err = json.Unmarshal(data, &users)
	return users, err
}

func SaveUsers(users []*entity.User) error {
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func GetUser(id int64) (*entity.User, error) {
	users, err := LoadUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Id == id {
			return user, nil
		}
	}
	return nil, fmt.Errorf("get user failed?")
}