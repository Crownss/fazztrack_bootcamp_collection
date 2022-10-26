package interfaces

import (
	"github.com/crownss/fazztrack_bootchamp/week_10/src/database/orm/models"
	"gorm.io/gorm"
)

type VehiclesRepo interface {
	FindAll() (*[]models.Vehicles, error)
	Save(data *models.Vehicles) (*models.Vehicles, error)
	Delete(name string) (*gorm.DB, error)
	Find(name string) (*models.Vehicles, error)
	Update(name string, data *models.Vehicles) (*models.Vehicles, error)
	MostPopular() (map[string]int, error)
	Sorting(str string) (*[]models.Vehicles, error)
	Search(str string) (*[]models.Vehicles, error)
	Category(str string) (*[]models.Vehicles, error)
}

type VehiclesService interface {
	GetAll() (*[]models.Vehicles, error)
	Add(data *models.Vehicles) (*models.Vehicles, error)
	Delete_Vehicles(name string) (*gorm.DB, error)
	Get(name string) (*models.Vehicles, error)
	Update_Vehicles(name string, data *models.Vehicles) (*models.Vehicles, error)
	MostPopular_Vehicles() (map[string]int, error)
	Sort_Vehicles(str string) (*[]models.Vehicles, error)
	Search_Vehicles(str string) (*[]models.Vehicles, error)
	Category_Vehicles(str string) (*[]models.Vehicles, error)
}
