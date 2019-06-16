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

	user, err := userService.Login(mobile, pwd)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespSuccess(writer, user, "Login Success")
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
