package users

import (
	"errors"

	"github.com/crownss/fazztrack_bootchamp/week_10/src/database/orm/models"
	"github.com/crownss/fazztrack_bootchamp/week_10/src/middleware"
	"golang.org/x/crypto/bcrypt"
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

func (user *users_repo) Find(username string) (*models.Users, error) {
	var data models.Users

	result := user.db.Where("username = ?", username).Find(&data).RowsAffected

	if result != 0 {
		return &data, nil

	}
	return nil, errors.New("Data not Found !")
}

func (user *users_repo) Delete(username string) (*gorm.DB, error) {
	var getid models.Users
	if err := user.db.Where("username = ?", username).First(&getid).Error; err != nil {
		return nil, err
	}
	deleting := user.db.Delete(&getid)
	return deleting, nil
}

func (user *users_repo) Update(username string, data *models.Users) (*models.Users, error) {
	var data_from_models models.Users
	result := user.db.Where("username = ?", username).First(&data_from_models).Updates(data)

	if result.Error != nil {
		return nil, errors.New("gagal menyimpan data")
	}

	return data, nil
}

func (user *users_repo) Register(request *models.RequestUsersRegister, data *models.Users) (*models.Users, error) {
	data.Username = request.Username
	query := user.db.Where("username = ?", request.Username).Find(&data)
	if query.RowsAffected != 0 {
		return nil, errors.New("username sudah digunakan")
	}
	if request.Password != request.Confirm_password {
		return nil, errors.New("confirm password harus sama dengan password")

	}
	encrypt, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	data.Password = string(encrypt)
	data.Is_Admin = request.Is_Admin
	save := user.db.Save(data)
	if save.Error != nil {
		return nil, errors.New("gagal menyimpan data")
	}
	return data, nil
}

func (user *users_repo) Login(request *models.RequestUsersLogin, data *models.Users) (string, error) {
	query := user.db.Where("username = ?", request.Username).Find(&data)
	if query.Error != nil {
		return "", errors.New("username tidak ditemukan")
	}
	if query.RowsAffected == 0 {
		return "", errors.New("username tidak ditemukan")
	}
	check := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(request.Password))
	if check != nil {
		return "", errors.New("password salah")
	}
	token := middleware.NewToken(request.Username, data.Is_Admin)
	create, err := token.Create()
	if err != nil {
		return err.Error(), err
	}
	return create, nil
}
