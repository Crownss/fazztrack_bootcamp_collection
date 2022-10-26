package models

import "gorm.io/gorm"

type History struct {
	gorm.Model
	UserID        string   `json:"user_id" form:"user_id" xml:"user_id"`
	User          Users    `gorm:"references:Username"`
	VehiclesID    string   `json:"vehicles_id" form:"vehicles_id" xml:"vehicles_id"`
	Vehicles      Vehicles `gorm:"references:Name"`
	StartRent     string   `json:"start_rent" form:"start_rent" xml:"start_rent"`
	EndRent       string   `json:"end_rent" form:"end_rent" xml:"end_rent"`
	PrePayment    uint     `json:"prepayment" form:"prepayment" xml:"prepayment"`
	PaymentStatus string   `json:"payment_status" form:"payment_status" xml:"payment_status"`
	ReturnStatus  string   `json:"return_status" form:"return_status" xml:"return_status"`
}