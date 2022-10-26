package vehicles

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/vehicles/").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", ctrl.GetAll).Methods("GET")
	route.HandleFunc("/q", ctrl.Sort_Vehicles).Methods("GET")
	route.HandleFunc("/search/{name}", ctrl.Search_Vehicles).Methods("GET")
	route.HandleFunc("/{id}", ctrl.Get).Methods("GET")
	route.HandleFunc("/popular/", ctrl.MostPopular_Vehicles).Methods("GET")
	route.HandleFunc("/", ctrl.Add).Methods("POST")
	route.HandleFunc("/{id}", ctrl.Update_Vehicles).Methods("PUT")
	route.HandleFunc("/{id}", ctrl.Delete_Vehicles).Methods("DELETE")
}
