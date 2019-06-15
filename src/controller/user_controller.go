package controller

import (
	"fmt"
	"math/rand"
	"model"
	"net/http"
	"service"
	"util"
)

var userService service.UserService

func UserLogin(writer http.ResponseWriter, request *http.Request) {

	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	pwd := request.PostForm.Get("passwd")

	loginOk := false
	if mobile == "15619258922" && pwd == "123456" {
		loginOk = true
	}

	if loginOk {
		data := make(map[string]interface{})
		data["id"] = 1
		data["token"] = "test"
		util.RespSuccess(writer, data, "")

	} else {
		util.RespFail(writer, "pwd is not correct!")
	}
}

func UserRegister(writer http.ResponseWriter, request *http.Request) {

	request.ParseForm()
	//
	mobile := request.PostForm.Get("mobile")
	//
	plainPwd := request.PostForm.Get("passwd")
	nickname := fmt.Sprintf("user%06d", rand.Int31())
	avatar := ""
	sex := model.SEX_UNKNOWN

	user, err := userService.Register(mobile, plainPwd, nickname, avatar, sex)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespSuccess(writer, user, "Register user successfully!")
	}
}
