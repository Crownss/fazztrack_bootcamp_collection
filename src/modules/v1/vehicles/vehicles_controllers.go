package vehicles

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/crownss/fazztrack_bootchamp/week_9/src/database/orm/models"
	"github.com/crownss/fazztrack_bootchamp/week_9/src/helpers"
	"github.com/crownss/fazztrack_bootchamp/week_9/src/interfaces"
	"github.com/gorilla/mux"
)

type vehicles_ctrl struct {
	service interfaces.VehiclesService
}

func NewCtrl(reps interfaces.VehiclesService) *vehicles_ctrl {
	return &vehicles_ctrl{reps}
}

func (vehicles *vehicles_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	data, err := vehicles.service.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helpers.Payload_400)
		return
	}
	payload := models.VehiclesResponseMany{
		Code:    http.StatusOK,
		Message: "data Found !",
		Status:  "success",
		Data:    *data,
	}
	json.NewEncoder(w).Encode(payload)
}

func (vehicles *vehicles_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var datas models.Vehicles
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helpers.Payload_400)
		return
	}

	_, err1 := vehicles.service.Add(&datas)
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helpers.Payload_400)
		return
	}
	json.NewEncoder(w).Encode(helpers.Payload_200)
}

func (vehicles *vehicles_ctrl) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	conv, _ := strconv.Atoi(id)
	data, err := vehicles.service.Get(conv)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helpers.Payload_400)
		return
	}
	payload := models.VehiclesResponseAny{
		Code:    http.StatusOK,
		Message: "data Found !",
		Status:  "success",
		Data:    *data,
	}
	json.NewEncoder(w).Encode(payload)
}

func (vehicles *vehicles_ctrl) MostPopular_Vehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	data, err := vehicles.service.MostPopular_Vehicles()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helpers.Payload_400)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (vehicles *vehicles_ctrl) Update_Vehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	conv, _ := strconv.Atoi(id)
	var datas models.Vehicles
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helpers.Payload_400)
		return
	}

	_, err1 := vehicles.service.Update_Vehicles(conv, &datas)
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helpers.Payload_400)
		return
	}
	json.NewEncoder(w).Encode(helpers.Payload_200)

}

func (vehicles *vehicles_ctrl) Delete_Vehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	conv, _ := strconv.Atoi(id)
	_, err := vehicles.service.Delete_Vehicles(conv)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helpers.Payload_400)
		return
	}
	json.NewEncoder(w).Encode(helpers.Payload_200)
}

func (vehicles *vehicles_ctrl) Sort_Vehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	sortby := r.URL.Query().Get("sortby")
	data, err := vehicles.service.Sort_Vehicles(sortby)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helpers.Payload_400)
		return
	}
	payload := models.VehiclesResponseMany{
		Code:    http.StatusOK,
		Message: "data Found !",
		Status:  http.StatusText(http.StatusOK),
		Data:    *data,
	}
	json.NewEncoder(w).Encode(payload)
}

func (vehicles *vehicles_ctrl) Search_Vehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	name := params["name"]
	data, err := vehicles.service.Search_Vehicles(strings.ToLower(name))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helpers.Payload_400)
		return
	}
	payload := models.VehiclesResponseMany{
		Code:    http.StatusOK,
		Message: "data Found !",
		Status:  http.StatusText(http.StatusOK),
		Data:    *data,
	}
	json.NewEncoder(w).Encode(payload)
}
