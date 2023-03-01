package student

import (
	"context"
	"log"
	"test_sat/model"
	"test_sat/pkg/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Insert(ctx context.Context, param model.Student) error {
	db, err := database.Connect(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, model.Student{
		Nama:  param.Nama,
		Umur:  param.Umur,
		Kelas: param.Kelas,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	return err
}

func Read(ctx context.Context) (*[]model.Student, error) {
	db, err := database.Connect(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	csr, err := db.Collection("student").Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	defer csr.Close(ctx)

	result := make([]model.Student, 0)
	for csr.Next(ctx) {
		var row model.Student
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}
	return &result, nil

}

func ReadDetail(ctx context.Context, id string) (*model.Student, error) {
	db, err := database.Connect(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	obId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	entry := &model.Student{}
	err = db.Collection("student").FindOne(ctx, bson.M{"_id": obId}).Decode(entry)
	if err != nil && err.Error() != "mongo: no documents in result" {
		log.Fatal(err.Error())
	}
	if entry.ID == "" {
		entry = nil
	}

	return entry, nil

}

func Update(ctx context.Context, id string, changes model.Student) error {
	db, err := database.Connect(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	obId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	var selector = bson.M{"_id": obId}
	_, err = db.Collection("student").UpdateOne(ctx, selector, bson.M{"$set": changes})
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

func Delete(ctx context.Context, id string) error {
	db, err := database.Connect(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	obId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	var selector = bson.M{"_id": obId}
	_, err = db.Collection("student").DeleteOne(ctx, selector)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}
