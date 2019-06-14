package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func UserLogin(writer http.ResponseWriter, request *http.Request) {

	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	pwd := request.PostForm.Get("password")

	loginOk := false
	if (mobile == "15619258922" && pwd == "123456") {
		loginOk = true
	}

	if loginOk {
		data := make(map[string]interface{})
		data["id"] = 1
		data["token"] = "test"
		Resp(writer, 0, data, "login successfully!")

	} else {
		Resp(writer, -1, nil, "pwd is not correct!")
	}
}

func Resp(w http.ResponseWriter, code int, data interface{}, msg string)  {

	rsp := Response{
		Code:code,
		Msg:msg,
		Data:data,
	}

	ret, err := json.Marshal(rsp)
	if err != nil {
		log.Printf("err %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([] byte(ret))
}

func main() {

	http.HandleFunc("/user/login", UserLogin)

	// start server
	http.ListenAndServe(":8080", nil)
}
