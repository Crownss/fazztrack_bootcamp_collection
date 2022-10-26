package users

import (
	"testing"

	"github.com/crownss/fazztrack_bootchamp/week_10/src/database/orm/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	repo := RepoMock{mock: mock.Mock{}}
	service := NewService(&repo)

	dataMock := []models.Users{
		{Username: "fantom12", Password: "password", Profile: "testing", Name: "fantom RHDB", Email: "fantom@gmail.com", Gender: "male", Address: "jalan...", Phone: "081532631588", Born: "212-98", Is_Admin: true}, {Username: "fantom13", Password: "password", Profile: "testing", Name: "fantom RHDB", Email: "fantom@gmail.com", Gender: "male", Address: "jalan...", Phone: "081532631588", Born: "212-98", Is_Admin: false}, {Username: "fantom14", Password: "password", Profile: "testing", Name: "fantom RHDB", Email: "fantom@gmail.com", Gender: "female", Address: "jalan...", Phone: "081532631588", Born: "212-98", Is_Admin: false},
	}
	repo.mock.On("FindAll").Return(&dataMock, nil)
	data, _ := service.GetAll()
	result := data
	for i, v := range *result {
		assert.Equal(t, dataMock[i].ID, v.ID, "id must be like data mock")
	}
}

func TestRegister(t *testing.T) {
	repo := RepoMock{mock: mock.Mock{}}
	service := NewService(&repo)

	dataMockRegister := models.RequestUsersRegister{
		Username: "fantom12", Password: "password", Confirm_password: "password",
	}
	dataMock := models.Users{
		Username: "fantom", Password: "password", Profile: "testing", Name: "fantom RHDB", Email: "fantom@gmail.com", Gender: "male", Address: "jalan...", Phone: "081532631588", Born: "212-98", Is_Admin: true,
	}

	repo.mock.On("Register", &dataMockRegister, &dataMock).Return(&dataMock, nil)
	data, err := service.Register_Users(&dataMockRegister, &dataMock)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	assert.NotEqual(t, dataMockRegister.Username, data.Username, "username must not equal")
}
