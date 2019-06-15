package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-xorm/xorm"
	"html/template"
	"log"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

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
		Resp(writer, 0, data, "login successfully!")

	} else {
		Resp(writer, -1, nil, "pwd is not correct!")
	}
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

func RegisterView() {
	tpl, err := template.ParseGlob("resources/view/**/*")
	if err != nil {
		log.Fatalf("parse global tpl error %s", err)
	}
	for _, v := range tpl.Templates() {
		tplName := v.Name()
		fmt.Println(tplName)
		http.HandleFunc(tplName, func(writer http.ResponseWriter, request *http.Request) {
			tpl.ExecuteTemplate(writer, tplName, nil)
		})
	}
}

var DbEngine *xorm.Engine

func init() {
	driverName := "mysql"
	dataSource := "root:agytorudhcv11@(localhost:3306)/golang_im?charset=utf8"
	DbEngine, err := xorm.NewEngine(driverName, dataSource)
	if err != nil {
		log.Fatalf("new engine err %s", err)
	}
	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(5)
}

func main() {

	http.HandleFunc("/user/login", UserLogin)
	RegisterView()
	// static resource handle
	http.Handle("/asset/", http.FileServer(http.Dir("./resources/")))
	// start server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("start server err %s", err)
	}
}
