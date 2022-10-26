package interfaces

import (
	"github.com/crownss/fazztrack_bootchamp/week_10/src/database/orm/models"
	"gorm.io/gorm"
)

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
