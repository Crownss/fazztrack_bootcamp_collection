package history

import (
	"errors"
	"strings"

	"github.com/crownss/fazztrack_bootchamp/week_10/src/database/orm/models"
	"gorm.io/gorm"
)

type history_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *history_repo {
	return &history_repo{db}
}

func (history *history_repo) FindAll() (*[]models.History, error) {
	var data []models.History

	result := history.db.Order("updated_at DESC").Find(&data)

	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}

	return &data, nil
}

func (history *history_repo) Save(data *models.History) (*models.History, error) {

	result := history.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("gagal menyimpan data")
	}

	return data, nil
}

func (history *history_repo) Find(str string) (*[]models.History, error) {
	var data []models.History

	result := history.db.Where("user_id = ?", str).Find(&data)

	if result.Error != nil {
		return nil, errors.New("data not found !")

	}
	return &data, nil
}

func (history *history_repo) Delete(id int) (*gorm.DB, error) {
	var getid models.History
	if err := history.db.Where("id = ?", id).First(&getid).Error; err != nil {
		return nil, err
	}
	deleting := history.db.Delete(&getid)
	return deleting, nil
}

func (history *history_repo) Update(id int, data *models.History) (*models.History, error) {
	var data_from_models models.History
	result := history.db.Where("id = ?", id).First(&data_from_models).Updates(data)

	if result.Error != nil {
		return nil, errors.New("gagal menyimpan data")
	}

	return data, nil
}

func (history *history_repo) Sorting(str string) (*[]models.History, error) {
	var data []models.History
	if strings.ToLower(str) == "desc" {
		result := history.db.Order("updated_at DESC").Find(&data)
		if result.Error != nil {
			return nil, errors.New("gagal mengambil data")
		}
		return &data, nil

	} else if strings.ToLower(str) == "asc" {
		result := history.db.Order("updated_at ASC").Find(&data)
		if result.Error != nil {
			return nil, errors.New("gagal mengambil data")
		}
		return &data, nil
	}
	return nil, errors.New("missing query")
}

func (history *history_repo) Search(str string) (*[]models.History, error) {
	var data []models.History
	result := history.db.Where("LOWER(vehicles_id) LIKE ?", "%"+str+"%").Or("LOWER(user_id) LIKE ?", "%"+str+"%").Find(&data)
	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}
	return &data, nil
}
