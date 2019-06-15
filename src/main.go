package main

import (
	"controller"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

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

func main() {

	http.HandleFunc("/user/login", controller.UserLogin)
	http.HandleFunc("/user/register", controller.UserRegister)
	RegisterView()
	// static resource handle
	http.Handle("/asset/", http.FileServer(http.Dir("./resources/")))
	// start server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("start server err %s", err)
	}
}
