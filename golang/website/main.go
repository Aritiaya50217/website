package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/Artitaya50217/website/golang/website/template"
	"github.com/gorilla/mux"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	_, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
func homePage(w http.ResponseWriter, r *http.Request) {
	// ใช้ฟังก์ชัน renderTemplate ให้แสดงไฟล์ index.html
	renderTemplate(w, "index")

}
func signupPage(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "signup")
}
func signinPage(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "signin")
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/signin", signinPage)
	myRouter.HandleFunc("/signup", signupPage)
	// port
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
func main() {
	handleRequests()

}
