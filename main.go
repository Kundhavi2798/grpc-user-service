package main

import (
	"fmt"
	"log"

	"github.com/Kundhavi2798/grpc-user-service/db"
	"github.com/Kundhavi2798/grpc-user-service/models"
	"github.com/Kundhavi2798/grpc-user-service/repository"
)

func main() {
	db.InitDB()
	id, err := repository.CreateUser(models.User{
		Name:  "Kundhavi",
		Email: "Kundhavi@gmail.com",
	})
	if err != nil {
		log.Fatal(err)
	}
	user, err := repository.GetUser(int32(id))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
}
