package routs

import (
	"github.com/gorilla/mux"
	"log"
	"master_go_programming/controllers/api"
	"net/http"
)

func Init() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/new-posts", api.CreateNewPosts).Methods("POST")


	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
