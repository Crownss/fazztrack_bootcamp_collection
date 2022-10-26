package history

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/history/").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", ctrl.GetAll).Methods("GET")
	route.HandleFunc("/{username}", ctrl.Get).Methods("GET")
	route.HandleFunc("/q", ctrl.Sort_History).Methods("GET")
	route.HandleFunc("/search/{obj}", ctrl.Search_History).Methods("GET")
	route.HandleFunc("/", ctrl.Add).Methods("POST")
	route.HandleFunc("/{id}", ctrl.Update_History).Methods("PUT")
	route.HandleFunc("/{id}", ctrl.Delete_History).Methods("DELETE")
}
