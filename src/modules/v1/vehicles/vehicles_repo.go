package vehicles

import (
	"errors"
	"strings"

	"github.com/crownss/fazztrack_bootchamp/week_10/src/database/orm/models"
	"gorm.io/gorm"
)

type vehicles_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *vehicles_repo {
	return &vehicles_repo{db}
}

func (vehicles *vehicles_repo) FindAll() (*[]models.Vehicles, error) {
	var data []models.Vehicles

	result := vehicles.db.Order("updated_at DESC").Find(&data).RowsAffected

	if result != 0 {
		return &data, nil
	}
	return nil, errors.New("gagal mengambil data")

}

func (vehicles *vehicles_repo) Category(str string) (*[]models.Vehicles, error) {
	var data []models.Vehicles

	result := vehicles.db.Where("category = ?",str).Order("updated_at DESC").Find(&data).RowsAffected

	if result != 0 {
		return &data, nil
	}
	return nil, errors.New("gagal mengambil data")

}

func (vehicles *vehicles_repo) Save(data *models.Vehicles) (*models.Vehicles, error) {

	result := vehicles.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("gagal menyimpan data")
	}

	return data, nil
}

func (vehicles *vehicles_repo) Find(name string) (*models.Vehicles, error) {
	var data models.Vehicles

	result := vehicles.db.Where("name = ?", name).First(&data).RowsAffected

	if result != 0 {
		return &data, nil

	}
	return nil, errors.New("Data not Found !")
}

func (vehicles *vehicles_repo) Delete(name string) (*gorm.DB, error) {
	var getid models.Vehicles
	if err := vehicles.db.Where("name = ?", name).First(&getid).Error; err != nil {
		return nil, err
	}
	deleting := vehicles.db.Delete(&getid)
	return deleting, nil
}

func (vehicles *vehicles_repo) Update(name string, data *models.Vehicles) (*models.Vehicles, error) {
	var data_from_models models.Vehicles
	result := vehicles.db.Where("name = ?", name).First(&data_from_models).Updates(data)

	if result.Error != nil {
		return nil, errors.New("gagal menyimpan data")
	}

	return data, nil
}

func (vehicles *vehicles_repo) MostPopular() (map[string]int, error) {
	var datas = new([]models.Popular_Vehicles)
	query := vehicles.db.Raw("select count(vehicles_id), name from histories join vehicles on histories.vehicles_id=vehicles.name group by name;").Scan(&datas)
	if query.Error != nil {
		return nil, errors.New("data not found !")

	}
	result := map[string]int{}
	for _, v := range *datas {
		result[v.Name] = v.Count
	}

	return result, nil
}

func (vehicles *vehicles_repo) Sorting(str string) (*[]models.Vehicles, error) {
	var data []models.Vehicles
	if strings.ToLower(str) == "desc" {
		result := vehicles.db.Order("updated_at DESC").Find(&data)
		if result.Error != nil {
			return nil, errors.New("gagal mengambil data")
		}
		return &data, nil

	} else if strings.ToLower(str) == "asc" {
		result := vehicles.db.Order("updated_at ASC").Find(&data)
		if result.Error != nil {
			return nil, errors.New("gagal mengambil data")
		}
		return &data, nil

	}
	return nil, errors.New("missing query")
}

func (vehicles *vehicles_repo) Search(str string) (*[]models.Vehicles, error) {
	var data []models.Vehicles
	result := vehicles.db.Where("LOWER(name) LIKE ?", "%"+str+"%").Find(&data)
	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}
	return &data, nil
}
