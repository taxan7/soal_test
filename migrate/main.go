package main

import (
	"context"
	"log"
	"test_sat/model"
	"test_sat/pkg/database"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Panic(err)
	}
	ctx := context.Background()
	db, err := database.Connect(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("user").InsertOne(ctx, model.User{
		Username: "admin",
		Password: "admin",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

}
