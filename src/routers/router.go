package routers

import (
	"errors"

	"github.com/crownss/fazztrack_bootchamp/week_8/src/database/orm"
	"github.com/crownss/fazztrack_bootchamp/week_8/src/modules/v1/history"
	"github.com/crownss/fazztrack_bootchamp/week_8/src/modules/v1/users"
	"github.com/crownss/fazztrack_bootchamp/week_8/src/modules/v1/vehicles"
	"github.com/gorilla/mux"
)

func New() (*mux.Router, error) {

	mainRoute := mux.NewRouter()

	db, err := orm.New()
	if err != nil {
		return nil, errors.New("gagal init database")
	}
	users.New(mainRoute, db)
	vehicles.New(mainRoute, db)
	history.New(mainRoute, db)

	return mainRoute, nil
}
