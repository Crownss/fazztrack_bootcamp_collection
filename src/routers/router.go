package routers

import (
	"errors"
	"net/http"

	"github.com/crownss/fazztrack_bootchamp/week_9/src/database/orm"
	"github.com/crownss/fazztrack_bootchamp/week_9/src/modules/v1/history"
	"github.com/crownss/fazztrack_bootchamp/week_9/src/modules/v1/users"
	"github.com/crownss/fazztrack_bootchamp/week_9/src/modules/v1/vehicles"
	"github.com/gorilla/mux"
)

func New() (*mux.Router, error) {

	mainRoute := mux.NewRouter()

	var dir string
	db, err := orm.New()
	if err != nil {
		return nil, errors.New("gagal init database")
	}
	mainRoute.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	users.New(mainRoute, db)
	vehicles.New(mainRoute, db)
	history.New(mainRoute, db)

	return mainRoute, nil
}
