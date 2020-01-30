package repository

import (
	"context"
	"log"

	"github.com/tarathep/hellogoapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//HelloLanguage is interface for implement
type HelloLanguage interface {
	AllHello() ([]*models.Hello, error)
	InsertHello(hello models.Hello) (models.Hello, error)
}

//AllHello get datal list
func (db *DB) AllHello() ([]*models.Hello, error) {

	collection := db.Database("test").Collection("hello")

	findOptions := options.Find()
	//findOptions.SetLimit(100)

	var results []*models.Hello

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		var elem models.Hello
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
	return results, nil
}

//InsertHello is insert data into mongo
func (db *DB) InsertHello(hello models.Hello) (models.Hello, error) {
	collection := db.Database("test").Collection("hello")
	_, err := collection.InsertOne(context.TODO(), hello)
	if err != nil {
		return hello, err
	}
	return hello, nil
}
