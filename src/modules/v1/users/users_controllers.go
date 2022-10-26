package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/crownss/fazztrack_bootchamp/week_8/src/database/orm/models"
	"github.com/crownss/fazztrack_bootchamp/week_8/src/helpers"
	"github.com/crownss/fazztrack_bootchamp/week_8/src/interfaces"
	"github.com/gorilla/mux"
)

type users_ctrl struct {
	service interfaces.UsersService
}

func NewCtrl(reps interfaces.UsersService) *users_ctrl {
	return &users_ctrl{reps}
}

func (user *users_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	data, err := user.service.GetAll()
	if err != nil {
		json.NewEncoder(w).Encode(helpers.Payload_400)
		return
	}
	payload := models.UsersResponseMany{
		Code:    http.StatusOK,
		Message: "data Found !",
		Status:  http.StatusText(http.StatusOK),
		Data:    *data,
	}
	json.NewEncoder(w).Encode(payload)
}

func (user *users_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var datas models.Users
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		json.NewEncoder(w).Encode(helpers.Payload_400)
		return
	}

	_, err1 := user.service.Add(&datas)
	if err1 != nil {
		json.NewEncoder(w).Encode(helpers.Payload_400)
		return
	}
	json.NewEncoder(w).Encode(helpers.Payload_200)

}

func (user *users_ctrl) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	conv, _ := strconv.Atoi(id)
	data, err := user.service.Get(conv)
	if err != nil {
		json.NewEncoder(w).Encode(helpers.Payload_400)
		return
	}
	payload := models.UsersResponseAny{
		Code:    http.StatusOK,
		Message: "data Found !",
		Status:  http.StatusText(http.StatusOK),
		Data:    *data,
	}
	json.NewEncoder(w).Encode(payload)
}

func (user *users_ctrl) Update_Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	conv, _ := strconv.Atoi(id)
	var datas models.Users
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		json.NewEncoder(w).Encode(helpers.Payload_400)
		return
	}

	_, err1 := user.service.Update_Users(conv, &datas)
	if err1 != nil {
		json.NewEncoder(w).Encode(helpers.Payload_400)
		return
	}
	json.NewEncoder(w).Encode(helpers.Payload_200)

}

func (user *users_ctrl) Delete_Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	conv, _ := strconv.Atoi(id)
	_, err := user.service.Delete_Users(conv)
	if err != nil {
		json.NewEncoder(w).Encode(helpers.Payload_400)
		return
	}
	json.NewEncoder(w).Encode(helpers.Payload_200)
}
