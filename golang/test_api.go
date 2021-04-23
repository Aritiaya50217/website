package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}
func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	// สร้างตัวแปร เชื่อมกับ router mux
	// r = http.Pequest
	vars := mux.Vars(r)
	// การค้นหาตามเลข id
	key := vars["id"]

	fmt.Fprintf(w, "Key: "+key)

	// Loop over all our Articles
	// if the article.Id equals the key we pass in
	// return the article encoded as JSON

	// จะออกจาก Loop เมื่อ id ตรงกับข้อมูลที่มี
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}

	}
}

// use method POST
// สร้าง Id ใหม่
func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string reponse containing the request body
	// append this to our Article array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	// แสดงข้อมูลแบบ json
	json.Unmarshal(reqBody, &article)
	// update our global Articles array to include
	// our new Article
	// append = การเพิ่มข้อมูลแบบ [] ลงในตัวแปร Articles ที่ประการแบบ global ไว้ด้านบน
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

// use method delete
func deleteArticle(w http.ResponseWriter, r *http.Request) {
	// once agian, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the 'id' of the article we
	// wish to delete
	id := vars["id"]
	// we then need to loop though all our articles
	for index, article := range Articles {
		// if our id path parameter matches one of our articles
		if article.Id == id {
			// update our Articles array to remove the article
			// เมื่อถูกลบออกไป ตำแหน่งถัดไปจะเลื่อนมาแทนที่
			Articles = append(Articles[:index], Articles[index+1])
		}

	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	// create a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	// Ordering is important here ! This has to be  defined before
	// the other '/article' endpoint
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	// add our new DELETE endpoint here
	myRouter.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "3", Title: "Hello 3", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
