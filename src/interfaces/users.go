package interfaces

import (
	"github.com/crownss/fazztrack_bootchamp/week_10/src/database/orm/models"
	"gorm.io/gorm"
)

type UsersRepo interface {
	FindAll() (*[]models.Users, error)
	Save(data *models.Users) (*models.Users, error)
	Delete(username string) (*gorm.DB, error)
	Find(username string) (*models.Users, error)
	Update(username string, data *models.Users) (*models.Users, error)
	Register(request *models.RequestUsersRegister, data *models.Users) (*models.Users, error)
	Login(request *models.RequestUsersLogin, data *models.Users) (string, error)
}

type UsersService interface {
	GetAll() (*[]models.Users, error)
	Add(data *models.Users) (*models.Users, error)
	Delete_Users(username string) (*gorm.DB, error)
	Get(username string) (*models.Users, error)
	Update_Users(username string, data *models.Users) (*models.Users, error)
	Register_Users(request *models.RequestUsersRegister, data *models.Users) (*models.Users, error)
	Login_Users(request *models.RequestUsersLogin, data *models.Users) (string, error)
}
