package libs

import (
	"encoding/json"
	"net/http"
)

type Respon struct {
	Code    int         `json:"code" form:"code" xml:"code"`
	Message interface{} `json:"message,omitempty" form:"message" xml:"message"`
	Status  string      `json:"status" form:"status" xml:"status"`
	IsError bool        `json:"isError" form:"isError" xml:"isError"`
	Data    interface{} `json:"data,omitempty" form:"data" xml:"data"`
}

func (res *Respon) Send(w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json")

	if res.IsError {
		w.WriteHeader(res.Code)
	}

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		w.Write([]byte("Error When Encode respone"))
	}
}

func Response(w http.ResponseWriter, data interface{}, code int, isError bool) *Respon {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	if isError {
		res := Respon{
			Code:    code,
			Message: data,
			Status:  http.StatusText(code),
			IsError: isError,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			w.Write([]byte("Error When Encode respone"))
		}
		return &res
	}
	res := Respon{
		Code:    code,
		Status:  http.StatusText(code),
		IsError: isError,
		Data:    data,
	}
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		w.Write([]byte("Error When Encode respone"))
	}
	return &res
}
