package users

import (
	"errors"

	"github.com/crownss/fazztrack_bootchamp/week_8/src/database/orm/models"
	"gorm.io/gorm"
)

type users_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *users_repo {
	return &users_repo{db}
}

func (user *users_repo) FindAll() (*[]models.Users, error) {
	var data []models.Users

	result := user.db.Order("updated_at DESC").Find(&data).RowsAffected

	if result != 0 {
		return &data, nil

	}
	return nil, errors.New("Data not Found !")
}



func (user *users_repo) Save(data *models.Users) (*models.Users, error) {

	result := user.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("gagal menyimpan data")
	}

	return data, nil
}

func (user *users_repo) Find(id int) (*models.Users, error) {
	var data models.Users

	result := user.db.Find(&data).RowsAffected

	if result != 0 {
		return &data, nil

	}
	return nil, errors.New("Data not Found !")
}

func (user *users_repo) Delete(id int) (*gorm.DB, error) {
	var getid models.Users
	if err := user.db.Where("id = ?", id).First(&getid).Error; err != nil {
		return nil, err
	}
	deleting := user.db.Delete(&getid)
	return deleting, nil
}

func (user *users_repo) Update(id int, data *models.Users)(*models.Users, error){
	var data_from_models models.Users
	result := user.db.Where("id = ?",id).First(&data_from_models).Updates(data)

	if result.Error != nil {
		return nil, errors.New("gagal menyimpan data")
	}

	return data, nil
}
