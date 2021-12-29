package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"master_go_programming/db"
	"master_go_programming/models"
	"net/http"
)

func CreateNewPosts(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var posts models.Post
	json.Unmarshal(reqBody, &posts)
	db.Init().Create(&posts)

	fmt.Println("Endpoint Hit: Creating New Booking")

	json.NewEncoder(w).Encode(posts)
}
