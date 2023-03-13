package infrastructure

import (
	"context"
	"time"

	"github.com/vnnyx/article-service/exception"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDatabase(configuration *Config) *mongo.Database {
	ctx, cancel := NewMongoContext()
	defer cancel()

	mongoPoolMin := configuration.MongoPoolMin
	mongoPoolMax := configuration.MongoPoolMax
	mongoMaxIdleTime := configuration.MongoMaxIdleTimeSecond

	option := options.Client().
		ApplyURI(configuration.MongoURI).
		SetMinPoolSize(uint64(mongoPoolMin)).
		SetMaxPoolSize(uint64(mongoPoolMax)).
		SetMaxConnIdleTime(time.Duration(mongoMaxIdleTime) * time.Second)

	client, err := mongo.NewClient(option)
	exception.PanicIfNeeded(err)

	err = client.Connect(ctx)
	exception.PanicIfNeeded(err)

	database := client.Database(configuration.MongoDatabae)
	return database
}

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
