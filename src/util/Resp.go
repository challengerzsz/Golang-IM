package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func RespFail(w http.ResponseWriter, msg string) {
	Resp(w, -1, nil, msg)
}

func RespSuccess(w http.ResponseWriter, data interface{}, msg string) {
	Resp(w, 0, data, msg)
}

func Resp(w http.ResponseWriter, code int, data interface{}, msg string) {

	rsp := Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}

	ret, err := json.Marshal(rsp)
	if err != nil {
		log.Printf("err %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(ret))
}
