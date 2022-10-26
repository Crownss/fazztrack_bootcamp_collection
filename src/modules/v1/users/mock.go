package users

import (
	"github.com/crownss/fazztrack_bootchamp/week_10/src/database/orm/models"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type RepoMock struct {
	mock mock.Mock
}

func (m *RepoMock) FindAll() (*[]models.Users, error) {
	args := m.mock.Called()
	return args.Get(0).(*[]models.Users), nil

}
func (m *RepoMock) Save(data *models.Users) (*models.Users, error) {
	args := m.mock.Called(data)
	return args.Get(0).(*models.Users), nil
}

func (m *RepoMock) Delete(username string) (*gorm.DB, error) {
	args := m.mock.Called(username)
	return args.Get(0).(*gorm.DB), nil
}
func (m *RepoMock) Find(username string) (*models.Users, error) {
	args := m.mock.Called(username)
	return args.Get(0).(*models.Users), nil
}
func (m *RepoMock) Update(username string, data *models.Users) (*models.Users, error) {
	args := m.mock.Called(username, data)
	return args.Get(0).(*models.Users), nil
}
func (m *RepoMock) Register(request *models.RequestUsersRegister, data *models.Users) (*models.Users, error) {
	args := m.mock.Called(request, data)
	return args.Get(0).(*models.Users), nil
}
func (m *RepoMock) Login(request *models.RequestUsersLogin, data *models.Users) (string, error) {
	args := m.mock.Called(request, data)
	return args.Get(0).(string), nil
}
