package history

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/crownss/fazztrack_bootchamp/week_10/src/database/orm/models"
	"github.com/crownss/fazztrack_bootchamp/week_10/src/interfaces"
	"github.com/crownss/fazztrack_bootchamp/week_10/src/libs"
	"github.com/gorilla/mux"
)

type history_ctrl struct {
	service interfaces.HistoryService
}

func NewCtrl(reps interfaces.HistoryService) *history_ctrl {
	return &history_ctrl{reps}
}

func (history *history_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := history.service.GetAll()
	if err != nil {
		libs.Response(w, err.Error(), 400, true)
		return
	}
	libs.Response(w, *data, 200, false)
}

func (history *history_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	var datas models.History
	err := json.NewDecoder(r.Body).Decode(&datas)
	println(err)
	if err != nil {
		libs.Response(w, err.Error(), 400, true)
		return
	}

	_, err1 := history.service.Add(&datas)
	if err1 != nil {
		libs.Response(w, err1.Error(), 400, true)
		return
	}
	libs.Response(w, "success create data!", 200, false)
}

func (history *history_ctrl) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]
	data, err := history.service.Get(username)
	if err != nil {
		libs.Response(w, err.Error(), 400, true)
		return
	}
	libs.Response(w, *data, 200, false)
}

func (history *history_ctrl) Update_History(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	conv, _ := strconv.Atoi(id)
	var datas models.History
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		libs.Response(w, err.Error(), 500, true)
		return
	}

	_, err1 := history.service.Update_History(conv, &datas)
	if err1 != nil {
		libs.Response(w, err.Error(), 400, true)
		return
	}
	libs.Response(w, "success update data!", 200, false)

}

func (history *history_ctrl) Delete_History(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	conv, _ := strconv.Atoi(id)
	_, err := history.service.Delete_History(conv)
	if err != nil {
		libs.Response(w, err.Error(), 500, true)
		return
	}
	libs.Response(w, "success delete data!", 200, false)
}

func (history *history_ctrl) Sort_History(w http.ResponseWriter, r *http.Request) {
	sortby := r.URL.Query().Get("sortby")
	data, err := history.service.Sort_History(sortby)
	println(sortby, data)
	if err != nil {
		libs.Response(w, err.Error(), 400, true)
		return
	}
	libs.Response(w, *data, 200, false)
}

func (history *history_ctrl) Search_History(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	search := params["obj"]
	data, err := history.service.Search_History(search)
	if err != nil {
		libs.Response(w, err.Error(), 400, true)
		return
	}
	libs.Response(w, *data, 200, false)
}
