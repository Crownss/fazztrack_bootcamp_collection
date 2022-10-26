package vehicles

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/crownss/fazztrack_bootchamp/week_10/src/database/orm/models"
	"github.com/crownss/fazztrack_bootchamp/week_10/src/interfaces"
	"github.com/crownss/fazztrack_bootchamp/week_10/src/libs"
	"github.com/gorilla/mux"
)

type vehicles_ctrl struct {
	service interfaces.VehiclesService
}

func NewCtrl(reps interfaces.VehiclesService) *vehicles_ctrl {
	return &vehicles_ctrl{reps}
}

func (vehicles *vehicles_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	category_query := r.URL.Query().Get("category")
	if category_query != "" {
		data, err := vehicles.service.Category_Vehicles(category_query)
		if err != nil {
			libs.Response(w, err.Error(), 400, true)
			return
		}
		libs.Response(w, &data, 200, false)
		return
	}
	data, err := vehicles.service.GetAll()
	if err != nil {
		libs.Response(w, err.Error(), 400, true)
		return
	}
	libs.Response(w, *data, 200, false)
}

func (vehicles *vehicles_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	var datas models.Vehicles
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		libs.Response(w, err.Error(), 400, true)
		return
	}

	data, err1 := vehicles.service.Add(&datas)
	if err1 != nil {
		libs.Response(w, err.Error(), 400, true)
		return
	}
	libs.Response(w, *data, 200, false)
}

func (vehicles *vehicles_ctrl) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	data, err := vehicles.service.Get(params["name"])
	if err != nil {
		libs.Response(w, err.Error(), 400, true)
		return
	}
	libs.Response(w, *data, 200, false)
}

func (vehicles *vehicles_ctrl) MostPopular_Vehicles(w http.ResponseWriter, r *http.Request) {
	data, err := vehicles.service.MostPopular_Vehicles()
	if err != nil {
		libs.Response(w, err.Error(), 400, true)
		return
	}
	libs.Response(w, data, 200, false)
}

func (vehicles *vehicles_ctrl) Update_Vehicles(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var datas models.Vehicles
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		libs.Response(w, err.Error(), 400, true)
		return
	}

	_, err1 := vehicles.service.Update_Vehicles(params["name"], &datas)
	if err1 != nil {
		libs.Response(w, err.Error(), 400, true)
		return
	}
	libs.Response(w, "success update data!", 200, false)

}

func (vehicles *vehicles_ctrl) Delete_Vehicles(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	_, err := vehicles.service.Delete_Vehicles(params["name"])
	if err != nil {
		libs.Response(w, err.Error(), 400, true)
		return
	}
	libs.Response(w, "success delete data!", 200, false)
}

func (vehicles *vehicles_ctrl) Sort_Vehicles(w http.ResponseWriter, r *http.Request) {
	sortby := r.URL.Query().Get("sortby")
	data, err := vehicles.service.Sort_Vehicles(sortby)
	if err != nil {
		libs.Response(w, err.Error(), 400, true)
		return
	}
	libs.Response(w, *data, 200, false)
}

func (vehicles *vehicles_ctrl) Search_Vehicles(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	data, err := vehicles.service.Search_Vehicles(strings.ToLower(name))
	if err != nil {
		libs.Response(w, err.Error(), 400, true)
		return
	}
	libs.Response(w, *data, 200, false)
}
