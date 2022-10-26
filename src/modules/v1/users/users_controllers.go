package users

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/crownss/fazztrack_bootchamp/week_10/src/database/orm/models"
	"github.com/crownss/fazztrack_bootchamp/week_10/src/interfaces"
	"github.com/crownss/fazztrack_bootchamp/week_10/src/libs"
	"github.com/gorilla/mux"
)

type users_ctrl struct {
	service interfaces.UsersService
}

func NewCtrl(reps interfaces.UsersService) *users_ctrl {
	return &users_ctrl{reps}
}

func (user *users_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := user.service.GetAll()
	if err != nil {
		libs.Response(w, err.Error(), 400, true)
		return
	}
	// check := r.Context().Value("is_admin").(bool)
	// if !check {
	// 	w.WriteHeader(http.StatusBadRequest)
	//	libs.Response("you're not admin", 403, true)
	// 	return
	// }
	libs.Response(w, *data, 200, false)
}

func (user *users_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	var datas models.Users
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		libs.Response(w, err.Error(), http.StatusBadRequest, true)
		return
	}

	_, err1 := user.service.Add(&datas)
	if err1 != nil {
		libs.Response(w, err.Error(), http.StatusBadRequest, true)
		return
	}
	libs.Response(w, "success post data!", 200, false)

}

func (user *users_ctrl) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]
	data, err := user.service.Get(username)
	if err != nil {
		libs.Response(w, err.Error(), 400, true)
		return
	}
	// payload := models.UsersResponseAny{
	// 	Code:    http.StatusOK,
	// 	Message: "data Found !",
	// 	Status:  http.StatusText(http.StatusOK),
	// 	Data:    *data,
	// }
	libs.Response(w, *data, 200, false)
}

func (user *users_ctrl) Update_Users(w http.ResponseWriter, r *http.Request) {
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
		libs.Response(w, err1.Error(), 400, true)

		return
	}
	libs.Response(w, "success update users!", 200, false)

}

func (user *users_ctrl) Delete_Users(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	_, err := user.service.Delete_Users(params["username"])
	if err != nil {
		libs.Response(w, err.Error(), 400, true)
		return
	}
	libs.Response(w, "success delete users!", 200, false)
}

func (user *users_ctrl) Register_Users(w http.ResponseWriter, r *http.Request) {
	var datas models.RequestUsersRegister
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		libs.Response(w, err.Error(), 400, true)
		return
	}

	_, err1 := user.service.Register_Users(&datas, &models.Users{})
	if err1 != nil {
		libs.Response(w, err.Error(), 400, true)
		return
	}
	libs.Response(w, "success register!", 200, false)
}

func (user *users_ctrl) Login_Users(w http.ResponseWriter, r *http.Request) {
	var datas *models.Users
	var req *models.RequestUsersLogin
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		libs.Response(w, err.Error(), 400, true)
		return
	}
	login, err1 := user.service.Login_Users(req, datas)
	if err1 != nil {
		libs.Response(w, err.Error(), 400, true)
		return
	}
	libs.Response(w, &login, 200, false)

}
