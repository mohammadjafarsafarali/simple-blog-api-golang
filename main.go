package main

import (
	_ "fmt"
	_ "github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"master_go_programming/routs"
)

func main() {

	//loading .env file
	err := godotenv.Load(".env")
	if err != nil {
		panic("There is a problem with env file!")
	}

	//pass requests to router handlers
	routs.Init()
}
