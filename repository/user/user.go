package user

import (
	"context"
	"log"
	user "test_sat/model"
	"test_sat/pkg/database"

	"go.mongodb.org/mongo-driver/bson"
)

func ReadUserPass(ctx context.Context, param user.User) (*user.User, error) {

	db, err := database.Connect(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	entry := &user.User{}
	err = db.Collection("user").FindOne(ctx, bson.M{"username": param.Username, "password": param.Password}).Decode(entry)
	if err != nil && err.Error() != "mongo: no documents in result" {
		log.Fatal(err.Error())
	}
	if entry.ID == "" {
		entry = nil
	}

	return entry, nil
}
