package users

import (
	"github.com/crownss/fazztrack_bootchamp/week_8/src/database/orm/models"
	"github.com/crownss/fazztrack_bootchamp/week_8/src/interfaces"
	"gorm.io/gorm"
)

type user_service struct {
	repo interfaces.UsersRepo
}

func NewService(repo interfaces.UsersRepo) *user_service {
	return &user_service{repo}
}

func (user *user_service) GetAll() (*[]models.Users, error) {
	data_users, err := user.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return data_users, nil
}

func (user *user_service) Add(data *models.Users) (*models.Users, error) {
	data, err := user.repo.Save(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (user *user_service) Delete_Users(id int)(*gorm.DB, error){
	data, err := user.repo.Delete(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (user *user_service) Get(id int) (*models.Users, error) {
	data_users, err := user.repo.Find(id)
	if err != nil {
		return nil, err
	}

	return data_users, nil
}

func (user *user_service) Update_Users(id int, data *models.Users) (*models.Users, error) {
	data, err := user.repo.Update(id, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}