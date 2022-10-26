package interfaces

import (
	"github.com/crownss/fazztrack_bootchamp/week_8/src/database/orm/models"
	"gorm.io/gorm"
)

type UsersRepo interface {
	FindAll() (*[]models.Users, error)
	Save(data *models.Users) (*models.Users, error)
	Delete(id int) (*gorm.DB, error)
	Find(id int) (*models.Users, error)
	Update(id int, data *models.Users) (*models.Users, error)
}

type UsersService interface {
	GetAll() (*[]models.Users, error)
	Add(data *models.Users) (*models.Users, error)
	Delete_Users(id int) (*gorm.DB, error)
	Get(id int) (*models.Users, error)
	Update_Users(id int, data *models.Users) (*models.Users, error)
}

type VehiclesRepo interface {
	FindAll() (*[]models.Vehicles, error)
	Save(data *models.Vehicles) (*models.Vehicles, error)
	Delete(id int) (*gorm.DB, error)
	Find(id int) (*models.Vehicles, error)
	Update(id int, data *models.Vehicles) (*models.Vehicles, error)
	MostPopular() (map[string]int, error)
	Sorting(str string) (*[]models.Vehicles, error)
	Search(str string) (*[]models.Vehicles, error)
}

type VehiclesService interface {
	GetAll() (*[]models.Vehicles, error)
	Add(data *models.Vehicles) (*models.Vehicles, error)
	Delete_Vehicles(id int) (*gorm.DB, error)
	Get(id int) (*models.Vehicles, error)
	Update_Vehicles(id int, data *models.Vehicles) (*models.Vehicles, error)
	MostPopular_Vehicles() (map[string]int, error)
	Sort_Vehicles(str string) (*[]models.Vehicles, error)
	Search_Vehicles(str string) (*[]models.Vehicles, error)
}

type HistoryRepo interface {
	FindAll() (*[]models.History, error)
	Save(data *models.History) (*models.History, error)
	Delete(id int) (*gorm.DB, error)
	Find(str string) (*[]models.History, error)
	Update(id int, data *models.History) (*models.History, error)
	Sorting(str string) (*[]models.History, error)
	Search(str string) (*[]models.History, error)
}

type HistoryService interface {
	GetAll() (*[]models.History, error)
	Add(data *models.History) (*models.History, error)
	Delete_History(id int) (*gorm.DB, error)
	Get(str string) (*[]models.History, error)
	Update_History(id int, data *models.History) (*models.History, error)
	Sort_History(str string) (*[]models.History, error)
	Search_History(str string) (*[]models.History, error)
}
