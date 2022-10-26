package interfaces

import (
	"github.com/crownss/fazztrack_bootchamp/week_9/src/database/orm/models"
	"gorm.io/gorm"
)

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
