package main

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"master_go_programming/db"
	"master_go_programming/models"
	"net/http"
)

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
	var posts models.Post
	json.Unmarshal(reqBody, &posts)
	db.Init().Create(&posts)

	fmt.Println("Endpoint Hit: Creating New Booking")
	json.NewEncoder(w).Encode(posts)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("There is a problem with env file!")
	}

	handleRequests()
}
