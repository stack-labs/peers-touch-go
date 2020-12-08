package util

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Msg  error  `json:"msg,omitempty"`
	Code string `json:"code,omitempty"`
}

type RspData struct {
	Error  *Error       `json:"error,omitempty"`
	Status int          `json:"status"`
	Data   *interface{} `json:"data,omitempty"`
}

func RspJsonErr(w http.ResponseWriter, err Error) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	rspData, _ := json.Marshal(RspData{
		Error:  &err,
		Status: 400,
	})

	_, _ = w.Write(rspData)
}

func RspJsonOK(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	rspData, _ := json.Marshal(RspData{
		Data:   &data,
		Status: 200,
	})

	_, _ = w.Write(rspData)
}
