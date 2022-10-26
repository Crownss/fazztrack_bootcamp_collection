package models

import "gorm.io/gorm"

type Vehicles struct {
	gorm.Model
	Img      string `json:"img" form:"img" xml:"img"`
	Name     string `json:"name" form:"name" xml:"name" gorm:"unique"`
	Location string `json:"location" form:"location" xml:"location"`
	Desc     string `json:"desc" form:"desc" xml:"desc"`
	Category string `json:"category" form:"category" xml:"category"`
	Like     uint   `json:"Like" form:"Like" xml:"Like"`
	Price    uint   `json:"price" form:"price" xml:"price"`
	Status   string `json:"status" form:"status" xml:"status"`
	Stock    uint   `json:"stock" form:"stock" xml:"stock"`
}

type Popular_Vehicles struct {
	Name  string `json:"name" form:"name" xml:"name"`
	Count int    `json:"count" form:"count" xml:"count"`
}
