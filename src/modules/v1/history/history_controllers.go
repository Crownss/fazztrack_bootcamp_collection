package history

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/crownss/fazztrack_bootchamp/week_9/src/database/orm/models"
	"github.com/crownss/fazztrack_bootchamp/week_9/src/helpers"
	"github.com/crownss/fazztrack_bootchamp/week_9/src/interfaces"
	"github.com/gorilla/mux"
)

type history_ctrl struct {
	service interfaces.HistoryService
}

func NewCtrl(reps interfaces.HistoryService) *history_ctrl {
	return &history_ctrl{reps}
}

func (history *history_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	data, err := history.service.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helpers.Payload_400)
		return
	}
	payload := models.HistoryResponseMany{
		Code:    http.StatusOK,
		Message: "data Found !",
		Status:  http.StatusText(http.StatusOK),
		Data:    *data,
	}
	json.NewEncoder(w).Encode(payload)
}

func (history *history_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var datas models.History
	err := json.NewDecoder(r.Body).Decode(&datas)
	println(err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helpers.Payload_400)
		return
	}

	_, err1 := history.service.Add(&datas)
	if err1 != nil {
		json.NewEncoder(w).Encode(helpers.Payload_400)
		return
	}
	json.NewEncoder(w).Encode(helpers.Payload_200)
}

func (history *history_ctrl) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	username := params["username"]
	data, err := history.service.Get(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helpers.Payload_400)
		return
	}
	payload := models.HistoryResponseMany{
		Code:    http.StatusOK,
		Message: "data Found !",
		Status:  http.StatusText(http.StatusOK),
		Data:    *data,
	}
	json.NewEncoder(w).Encode(payload)
}

func (history *history_ctrl) Update_History(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	conv, _ := strconv.Atoi(id)
	var datas models.History
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helpers.Payload_500)
		return
	}

	_, err1 := history.service.Update_History(conv, &datas)
	if err1 != nil {
		json.NewEncoder(w).Encode(helpers.Payload_400)
		return
	}
	json.NewEncoder(w).Encode(helpers.Payload_200)

}

func (history *history_ctrl) Delete_History(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	conv, _ := strconv.Atoi(id)
	_, err := history.service.Delete_History(conv)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helpers.Payload_500)
		return
	}
	json.NewEncoder(w).Encode(helpers.Payload_200)
}

func (history *history_ctrl) Sort_History(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	sortby := r.URL.Query().Get("sortby")
	data, err := history.service.Sort_History(sortby)
	println(sortby, data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helpers.Payload_400)
		return
	}
	payload := models.HistoryResponseMany{
		Code:    http.StatusOK,
		Message: "data Found !",
		Status:  http.StatusText(http.StatusOK),
		Data:    *data,
	}
	json.NewEncoder(w).Encode(payload)
}

func (history *history_ctrl) Search_History(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	search := params["obj"]
	data, err := history.service.Search_History(search)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helpers.Payload_400)
		return
	}
	payload := models.HistoryResponseMany{
		Code:    http.StatusOK,
		Message: "data Found !",
		Status:  http.StatusText(http.StatusOK),
		Data:    *data,
	}
	json.NewEncoder(w).Encode(payload)
}
