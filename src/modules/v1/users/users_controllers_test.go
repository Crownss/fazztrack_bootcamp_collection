package users

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/crownss/fazztrack_bootchamp/week_10/src/database/orm/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repo = RepoMock{mock.Mock{}}
var service = NewService(&repo)
var ctrl = NewCtrl(service)

var dataMock = []models.Users{
	{Username: "fantom12", Password: "password", Profile: "testing", Name: "fantom RHDB", Email: "fantom@gmail.com", Gender: "male", Address: "jalan...", Phone: "081532631588", Born: "212-98", Is_Admin: true}, {Username: "fantom13", Password: "password", Profile: "testing", Name: "fantom RHDB", Email: "fantom@gmail.com", Gender: "male", Address: "jalan...", Phone: "081532631588", Born: "212-98", Is_Admin: false}, {Username: "fantom14", Password: "password", Profile: "testing", Name: "fantom RHDB", Email: "fantom@gmail.com", Gender: "female", Address: "jalan...", Phone: "081532631588", Born: "212-98", Is_Admin: false},
}

func TestCtrlGetAll(t *testing.T) {
	w := httptest.NewRecorder()
	mux := mux.NewRouter()

	repo.mock.On("FindAll").Return(&dataMock, nil)
	mux.HandleFunc("/testing/getall", ctrl.GetAll)
	mux.ServeHTTP(w, httptest.NewRequest("GET", "/testing/getall", nil))
	response := models.UsersResponseMany{
		Code:    200,
		Message: "success !",
		Status:  "success !",
		Data:    dataMock,
	}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal(err.Error())
		return
	}
	fmt.Println(response)
	assert.Equal(t, response.Code, w.Code, "Code must be 200")
}
