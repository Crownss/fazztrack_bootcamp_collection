package history

import (
	"github.com/crownss/fazztrack_bootchamp/week_10/src/database/orm/models"
	"github.com/crownss/fazztrack_bootchamp/week_10/src/interfaces"
	"gorm.io/gorm"
)

type history_service struct {
	repo interfaces.HistoryRepo
}

func NewService(repo interfaces.HistoryRepo) *history_service {
	return &history_service{repo}
}

func (history *history_service) GetAll() (*[]models.History, error) {
	data, err := history.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (history *history_service) Add(data *models.History) (*models.History, error) {
	data, err := history.repo.Save(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (history *history_service) Delete_History(id int) (*gorm.DB, error) {
	data, err := history.repo.Delete(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (history *history_service) Get(str string) (*[]models.History, error) {
	data_history, err := history.repo.Find(str)
	if err != nil {
		return nil, err
	}

	return data_history, nil
}

func (history *history_service) Update_History(id int, data *models.History) (*models.History, error) {
	data, err := history.repo.Update(id, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (history *history_service) Sort_History(str string) (*[]models.History, error) {
	data, err := history.repo.Sorting(str)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (history *history_service) Search_History(str string) (*[]models.History, error) {
	data, err := history.repo.Search(str)
	if err != nil {
		return nil, err
	}
	return data, nil
}
