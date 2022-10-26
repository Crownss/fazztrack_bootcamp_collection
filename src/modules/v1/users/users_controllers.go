package users

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/crownss/fazztrack_bootchamp/week_9/src/database/orm/models"
	"github.com/crownss/fazztrack_bootchamp/week_9/src/helpers"
	"github.com/crownss/fazztrack_bootchamp/week_9/src/interfaces"
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
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Status:  http.StatusText(http.StatusBadRequest),
		})
		return
	}
	// check := r.Context().Value("is_admin").(bool)
	// if !check {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	json.NewEncoder(w).Encode(models.Response{
	// 		Code:    http.StatusBadRequest,
	// 		Message: "you are not admin !",
	// 		Status:  http.StatusText(http.StatusBadRequest),
	// 	})
	// 	return
	// }
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
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Status:  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	_, err1 := user.service.Add(&datas)
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Status:  http.StatusText(http.StatusBadRequest),
		})
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
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Status:  http.StatusText(http.StatusBadRequest),
		})
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
	// params := mux.Vars(r)
	// id := params["id"]
	// conv, _ := strconv.Atoi(id)
	var datas models.Users
	// err := json.NewDecoder(r.Body).Decode(&datas)
	// if err != nil {
	// 	json.NewEncoder(w).Encode(helpers.Payload_400)
	// 	return
	// }
	// err := r.FormValue(r.Body)
	datas.Profile = "http://0.0.0.0/" + os.Getenv("PORT") + "/static/" + r.Context().Value("filename").(string)
	_, err1 := user.service.Update_Users(r.Context().Value("username").(string), &datas)
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Code:    http.StatusBadRequest,
			Message: err1.Error(),
			Status:  http.StatusText(http.StatusBadRequest),
		})
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
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Status:  http.StatusText(http.StatusBadRequest),
		})
		return
	}
	json.NewEncoder(w).Encode(helpers.Payload_200)
}

func (user *users_ctrl) Register_Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var datas models.RequestUsersRegister
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Status:  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	_, err1 := user.service.Register_Users(&datas, &models.Users{})
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Status:  http.StatusText(http.StatusBadRequest),
		})
		return
	}
	json.NewEncoder(w).Encode(helpers.Payload_200)
}

func (user *users_ctrl) Login_Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var datas *models.Users
	var req *models.RequestUsersLogin
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Status:  http.StatusText(http.StatusBadRequest),
		})
		return
	}
	login, err1 := user.service.Login_Users(req, datas)
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Status:  http.StatusText(http.StatusBadRequest),
		})
		return
	}
	paylaod := &models.Token_Response{
		Code:    http.StatusOK,
		Message: "Login success !",
		Status:  http.StatusText(http.StatusOK),
		Token:   login,
	}
	json.NewEncoder(w).Encode(paylaod)

}
