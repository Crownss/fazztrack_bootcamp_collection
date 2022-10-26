package users

import (
	"github.com/crownss/fazztrack_bootchamp/week_9/src/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/users/").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", middleware.CheckAdmin(ctrl.GetAll)).Methods("GET")
	route.HandleFunc("/{id}", ctrl.Get).Methods("GET")
	route.HandleFunc("/register", ctrl.Register_Users).Methods("POST")
	route.HandleFunc("/login", ctrl.Login_Users).Methods("POST")
	route.HandleFunc("/update", middleware.MultipleMiddleware(ctrl.Update_Users, middleware.CheckAuth, middleware.UploadProfile)).Methods("PUT")
	route.HandleFunc("/{id}", ctrl.Delete_Users).Methods("DELETE")
}
