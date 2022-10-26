package vehicles

import (
	"github.com/crownss/fazztrack_bootchamp/week_10/src/database/orm/models"
	"github.com/crownss/fazztrack_bootchamp/week_10/src/interfaces"
	"gorm.io/gorm"
)

type vehicles_service struct {
	repo interfaces.VehiclesRepo
}

func NewService(repo interfaces.VehiclesRepo) *vehicles_service {
	return &vehicles_service{repo}
}

func (vehicles *vehicles_service) GetAll() (*[]models.Vehicles, error) {
	data, err := vehicles.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (vehicles *vehicles_service) Add(data *models.Vehicles) (*models.Vehicles, error) {
	data, err := vehicles.repo.Save(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (vehicles *vehicles_service) Delete_Vehicles(name string) (*gorm.DB, error) {
	data, err := vehicles.repo.Delete(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (vehicles *vehicles_service) Get(name string) (*models.Vehicles, error) {
	data_vehicles, err := vehicles.repo.Find(name)
	if err != nil {
		return nil, err
	}

	return data_vehicles, nil
}

func (vehicles *vehicles_service) Update_Vehicles(name string, data *models.Vehicles) (*models.Vehicles, error) {
	data, err := vehicles.repo.Update(name, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (vehicles *vehicles_service) MostPopular_Vehicles() (map[string]int, error) {
	data, err := vehicles.repo.MostPopular()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (vehicles *vehicles_service) Sort_Vehicles(str string) (*[]models.Vehicles, error) {
	data, err := vehicles.repo.Sorting(str)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (vehicles *vehicles_service) Search_Vehicles(str string) (*[]models.Vehicles, error) {
	data, err := vehicles.repo.Search(str)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (vehicles *vehicles_service)Category_Vehicles(str string)(*[]models.Vehicles, error){
	data, err := vehicles.repo.Category(str)
	if err != nil{
		return nil, err
	}
	return data, nil
}
