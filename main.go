package main

import (
	"html/template"
	"net/http"
	"log"
	"strings"
)
var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}
func main(){
	http.HandleFunc("/index",idx)
	http.HandleFunc("/about",abot)
	http.HandleFunc("/contact",cntct)
	http.HandleFunc("/apply",aply)
	http.HandleFunc("/index/ajab/",sayHello)

	err := http.ListenAndServe(":8080",nil)
	if err != nil{
		panic(err)
	}
}


func idx(w http.ResponseWriter, req *http.Request){
	err := tpl.ExecuteTemplate(w,"index.gohtml",nil)
	if err != nil{
		log.Println(err)
		http.Error(w,"Internal server error",http.StatusInternalServerError)
	}
}

func abot(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	if err != nil{
		log.Println(err)
		http.Error(w,"Internal server error",http.StatusInternalServerError)
	}
}

func cntct(w http.ResponseWriter, req *http.Request){
	err := tpl.ExecuteTemplate(w,"contact.gohtml",nil)
	if err != nil{
		log.Println(err)
		http.Error(w,"Internal server error",http.StatusInternalServerError)
	}
}

func aply(w http.ResponseWriter, req *http.Request){
	err := tpl.ExecuteTemplate(w,"apply.gohtml",nil)
	if err != nil{
		log.Println(err)
		http.Error(w,"Internal server error",http.StatusInternalServerError)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/index/ajab/")
	message = "Hello " + message
	w.Write([]byte(message))
}