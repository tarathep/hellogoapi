package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DB is mongodb
type DB struct {
	*mongo.Client
}

//Init is Init database mongo
func Init(dataSourceName string) (*DB, error) {
	clientOptions := options.Client().ApplyURI(dataSourceName)

	db, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
