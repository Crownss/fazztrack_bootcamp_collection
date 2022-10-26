package users

import (
	"github.com/crownss/fazztrack_bootchamp/week_10/src/database/orm/models"
	"github.com/crownss/fazztrack_bootchamp/week_10/src/interfaces"
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

func (user *user_service) Delete_Users(username string) (*gorm.DB, error) {
	data, err := user.repo.Delete(username)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (user *user_service) Get(username string) (*models.Users, error) {
	data_users, err := user.repo.Find(username)
	if err != nil {
		return nil, err
	}

	return data_users, nil
}

func (user *user_service) Update_Users(username string, data *models.Users) (*models.Users, error) {
	data, err := user.repo.Update(username, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (user *user_service) Register_Users(request *models.RequestUsersRegister, data *models.Users) (*models.Users, error) {
	datas, err := user.repo.Register(request, data)
	if err != nil {
		return nil, err
	}

	return datas, nil
}

func (user *user_service) Login_Users(request *models.RequestUsersLogin, data *models.Users) (string, error) {
	datas, err := user.repo.Login(request, data)
	if err != nil {
		return "", err
	}

	return datas, nil
}
