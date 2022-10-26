package interfaces

import (
	"github.com/crownss/fazztrack_bootchamp/week_9/src/database/orm/models"
	"gorm.io/gorm"
)

type UsersRepo interface {
	FindAll() (*[]models.Users, error)
	Save(data *models.Users) (*models.Users, error)
	Delete(id int) (*gorm.DB, error)
	Find(id int) (*models.Users, error)
	Update(username string, data *models.Users) (*models.Users, error)
	Register(request *models.RequestUsersRegister, data *models.Users) (*models.Users, error)
	Login(request *models.RequestUsersLogin, data *models.Users) (string, error)
}

type UsersService interface {
	GetAll() (*[]models.Users, error)
	Add(data *models.Users) (*models.Users, error)
	Delete_Users(id int) (*gorm.DB, error)
	Get(id int) (*models.Users, error)
	Update_Users(username string, data *models.Users) (*models.Users, error)
	Register_Users(request *models.RequestUsersRegister, data *models.Users) (*models.Users, error)
	Login_Users(request *models.RequestUsersLogin, data *models.Users) (string, error)
}
