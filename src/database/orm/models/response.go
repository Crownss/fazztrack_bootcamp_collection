package models

type VehiclesResponseAny struct {
	Code    int      `json:"code" form:"code" xml:"code"`
	Message string   `json:"message" form:"message" xml:"message"`
	Status  string   `json:"status" form:"status" xml:"status"`
	Data    Vehicles `json:"data" form:"data" xml:"data"`
}

type VehiclesResponseMany struct {
	Code    int        `json:"code" form:"code" xml:"code"`
	Message string     `json:"message" form:"message" xml:"message"`
	Status  string     `json:"status" form:"status" xml:"status"`
	Data    []Vehicles `json:"data" form:"data" xml:"data"`
}

type UsersResponseAny struct {
	Code    int    `json:"code" form:"code" xml:"code"`
	Message string `json:"message" form:"message" xml:"message"`
	Status  string `json:"status" form:"status" xml:"status"`
	Data    Users  `json:"data" form:"data" xml:"data"`
}

type UsersResponseMany struct {
	Code    int     `json:"code" form:"code" xml:"code"`
	Message string  `json:"message" form:"message" xml:"message"`
	Status  string  `json:"status" form:"status" xml:"status"`
	Data    []Users `json:"data" form:"data" xml:"data"`
}

type HistoryResponseMany struct {
	Code    int       `json:"code" form:"code" xml:"code"`
	Message string    `json:"message" form:"message" xml:"message"`
	Status  string    `json:"status" form:"status" xml:"status"`
	Data    []History `json:"data" form:"data" xml:"data"`
}

type HistoryResponseAny struct {
	Code    int     `json:"code" form:"code" xml:"code"`
	Message string  `json:"message" form:"message" xml:"message"`
	Status  string  `json:"status" form:"status" xml:"status"`
	Data    History `json:"data" form:"data" xml:"data"`
}
