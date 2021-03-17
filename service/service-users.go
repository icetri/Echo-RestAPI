package service

import (
	"echo-restapi/types"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
)

type Service struct {
	max   int
	users *types.Users
}

func NewService() *Service {
	u := &types.Users{}
	return &Service{
		users: u,
	}
}

func (s *Service) SaveUser(user *types.User, oldusers []types.User) ([]types.User, error) {
	oldusers = append(oldusers, *user)
	err := UpdateFileJSON(oldusers, "json.json")
	if err != nil {
		return nil, errors.Wrap(err, "Ошибка сервиса UpdateFileJSON в SaveUser")
	}
	return oldusers, nil
}

func (s *Service) GetUser(id int) (*types.User, error) {
	tylpe := types.User{}
	b, err := ioutil.ReadFile("json.json")
	if err != nil {
		return nil, fmt.Errorf("Пользователей нет")
	}
	err = json.Unmarshal(b, &s.users.Users)
	if err != nil {
		return nil, errors.Wrap(err, "Ошибка декодирования сервиса GetUser")
	}
	for _, val := range s.users.Users {
		if val.Id == id {
			err := json.Unmarshal(b, &s.users.Users)
			if err != nil {
				return nil, errors.Wrap(err, "Ошибка декодирования Service GetUser")
			}
			tylpe = val
		}
	}
	return &tylpe, err
}

func (s *Service) GetUsers() ([]types.User, error) {
	b, err := ioutil.ReadFile("json.json")
	if err != nil {
		return nil, fmt.Errorf("Пользователей нет")
	}
	x := s.users.Users
	err = json.Unmarshal(b, &x)
	if err != nil {
		return nil, errors.Wrap(err, "Ошибка декодирования GetUsers")
	}
	return x, err
}

func (s *Service) UpdateUser(id int, name string) (bool, error) {
	b, err := ioutil.ReadFile("json.json")
	if err != nil {
		return true, fmt.Errorf("Пользователей нет")
	}
	err = json.Unmarshal(b, &s.users.Users)
	if err != nil {
		return true, errors.Wrap(err, "Ошибка декодирования сервиса UpdateUser")
	}
	okId := true
	for i, val := range s.users.Users {
		if id != val.Id {
			okId = false
			continue
		}
		if val.Id == id {
			val.Name = name
			s.users.Users[i].Name = val.Name
			okId = true
			break
		}
	}
	err = UpdateFileJSON(s.users.Users, "json.json")
	if err != nil {
		return true, errors.Wrap(err, "Ошибка сервиса UpdateFileJSON в UpdateUser")
	}
	return okId, err
}

func (s *Service) DeleteUser(id int) (bool, error) {
	b, err := ioutil.ReadFile("json.json")
	if err != nil {
		return true, fmt.Errorf("Пользователей нет")
	}
	err = json.Unmarshal(b, &s.users.Users)
	if err != nil {
		return true, errors.Wrap(err, "Ошибка декодирования DeleteUser")
	}
	okId := true
	for i, val := range s.users.Users {
		if id != val.Id {
			okId = false
			continue
		}
		if id == val.Id {
			s.users.Users = append(s.users.Users[:i], s.users.Users[i+1:]...)
			okId = true
			break
		}
	}
	err = UpdateFileJSON(s.users.Users, "json.json")
	if err != nil {
		return true, errors.Wrap(err, "Ошибка сервиса UpdateFileJSON в DeleteUser")
	}
	return okId, err
}
