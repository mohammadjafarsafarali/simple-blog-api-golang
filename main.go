package main

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"log"
	"net/http"
)

var db *gorm.DB
var err error

type Posts struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string    `json:"content"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to HomePage!")
	fmt.Println("Endpoint Hit: HomePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/new-posts", createNewPosts).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}


func createNewPosts(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var posts Posts
	json.Unmarshal(reqBody, &posts)
	db.Create(&posts)
	fmt.Println("Endpoint Hit: Creating New Booking")
	json.NewEncoder(w).Encode(posts)
}



func main() {
	// defining username and password and mysql config
	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/blog_go?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}

	db.AutoMigrate(&Posts{})

	handleRequests()
}
