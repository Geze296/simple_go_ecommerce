package database

import (
	"context"
	"time"
	"github.com/geze296/simple_go_ecommerce/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongo(config *config.MongoConfig) (*mongo.Client, error){

	ctx, cancel := context.WithTimeout(context.Background(),10*time.Second)

	defer cancel()
	
	client,err := mongo.Connect(ctx,options.Client().ApplyURI(config.MongoURL))

	if err !=nil{
		return nil,err
	}

	return client,nil
}